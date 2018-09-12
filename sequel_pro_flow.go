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

func NewSequelProFlow(config *SequelProDatasources) (spf *SequelProFlow, err error) {
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
		SequelProDatasources: *config,
		plist:                sequelPlist,
	}, nil
}

type SequelProFlow struct {
	SequelProDatasources
	plist       SequelProRootPlist
	credentials *Credentials
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

func (sp *SequelProFlow) Name() string {
	return "SEQUEL PRO"
}

func (sp *SequelProFlow) UpdateWithCredentials(credentials *Credentials) (err error) {
	plistItem, found := sp.findPlistItem()
	if !found {
		return errors.New(fmt.Sprintf("[SEQUEL PRO] Could not find plist item for database UUID %d", sp.DatabaseUUID))
	}

	plistItem.User = credentials.Username
	sp.credentials = credentials

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
	log.Printf("[SEQUEL PRO] Persisting username %s and password %s to service %s and account %s", plistItem.User, sp.credentials.Password, service, account)

	bytes, err := plist.Marshal(sp.plist, SEQUEL_PRO_PLIST_FORMAT)
	if err != nil {
		return
	}

	ioutil.WriteFile(sp.SequelProPlistPath, bytes, 0644)

	return CreateOrUpdateKeychainEntriesForService(service, account, sp.credentials.Password, []string{})
}
