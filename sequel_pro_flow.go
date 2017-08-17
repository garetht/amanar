package main

import (
	"io/ioutil"

	"fmt"
	"log"

	"errors"

	"github.com/DHowett/go-plist"
)

const SEQUEL_PRO_PLIST_FORMAT = plist.XMLFormat

type SequelProRootPlist struct {
	FavoritesRoot struct {
		IsExpanded bool                 `plist:"IsExpanded"`
		Name       string               `plist:"Name"`
		Children   []SequelProPlistItem `plist:"Children"`
	} `plist:"Favorites Root"`
}

type SequelProPlistItem struct {
	ColorIndex                        int64  `plist:"colorIndex"`
	Database                          string `plist:"database"`
	Host                              string `plist:"host"`
	Id                                int64  `plist:"id"`
	Name                              string `plist:"name"`
	Port                              string `plist:"port"`
	Socket                            string `plist:"socket"`
	SSHHost                           string `plist:"sshHost"`
	SSHKeyLocation                    string `plist:"sshKeyLocation"`
	SSHKeyLocationEnabled             int    `plist:"sshKeyLocationEnabled"`
	SSHPort                           string `plist:"sshPort"`
	SSHUser                           string `plist:"sshUser"`
	SSLCACertFileLocation             string `plist:"sslcaCertFileLocation"`
	SSLCACertFileLocationEnabled      int    `plist:"sslcaCertFileLocationEnabled"`
	SSLCertificateFileLocation        string `plist:"sslCertificateFileLocation"`
	SSLCertificateFileLocationEnabled int    `plist:"sslCertificateFileLocationEnabled"`
	SSLKeyFileLocation                string `plist:"sslKeyFileLocation"`
	SSLKeyFileLocationEnabled         int    `plist:"sslKeyFileLocationEnabled"`
	Type                              int    `plist:"type"`
	UseSSL                            int    `plist:"useSSL"`
	User                              string `plist:"user"`
}

func NewSequelProFlow(config *SequelProDatasourcesConfig) (spf *SequelProFlow, err error) {
	bytes, err := ioutil.ReadFile(config.SequelProPlistPath)
	if err != nil {
		return
	}

	sequelPlist := SequelProRootPlist{}
	_, err = plist.Unmarshal(bytes, &sequelPlist)
	if err != nil {
		return
	}

	return &SequelProFlow{
		SequelProDatasourcesConfig: *config,
		plist: sequelPlist,
	}, nil
}

type SequelProFlow struct {
	SequelProDatasourcesConfig
	plist            SequelProRootPlist
	passwordToUpdate string
}

func (sp *SequelProFlow) findPlistItem() (plistItem *SequelProPlistItem, foundItem bool) {
	children := sp.plist.FavoritesRoot.Children
	// N.B. the iteratees of range will be copied values. This allows
	// us to refer to the indexed children as actual pointers
	for i := 0; i < len(children); i++ {
		if fmt.Sprintf("%d", children[i].Id) == sp.DatabaseUUID {
			plistItem = &children[i]
			foundItem = true
			break
		}
	}
	return
}

func (sp *SequelProFlow) UpdateUsername(username string) (err error) {
	plistItem, found := sp.findPlistItem()
	if !found {
		return errors.New(fmt.Sprintf("[SEQUEL PRO] Could not find plist item for database UUID %d", sp.DatabaseUUID))
	}

	plistItem.User = username
	return
}

func (sp *SequelProFlow) UpdatePassword(password string) error {
	sp.passwordToUpdate = password
	return nil
}

func (sp *SequelProFlow) PersistChanges() (err error) {
	plistItem, found := sp.findPlistItem()
	if !found {
		return errors.New("Could not find a matching Sequel Database for that Sequel UUID.")
	}

	// These two values need to be synchronized for Sequel Pro to be able to
	// read the correct keychain value.
	service := fmt.Sprintf("Sequel Pro : %s (%d)", plistItem.Name, plistItem.Id)
	account := fmt.Sprintf("%s@%s/%s", plistItem.User, plistItem.Host, plistItem.Database)
	log.Printf("[SEQUEL PRO] Persisting username %s and password %s to service %s and account %s", plistItem.User, sp.passwordToUpdate, service, account)

	bytes, err := plist.Marshal(sp.plist, SEQUEL_PRO_PLIST_FORMAT)
	if err != nil {
		return
	}

	ioutil.WriteFile(sp.SequelProPlistPath, bytes, 0644)

	return CreateOrUpdateKeychainEntriesForService(service, account, sp.passwordToUpdate, []string{})
}

func (sp *SequelProFlow) UpdateCredentials(credentials *Credentials) (err error) {
	err = sp.UpdateUsername(credentials.Username)
	if err != nil {
		return
	}

	err = sp.UpdatePassword(credentials.Password)
	if err != nil {
		return
	}

	err = sp.PersistChanges()
	if err != nil {
		return
	}

	return nil
}
