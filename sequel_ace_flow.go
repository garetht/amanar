package amanar

import (
	"io/ioutil"

	"fmt"
	"log"

	"errors"

	"howett.net/plist"
)

const SEQUEL_ACE_PLIST_FORMAT = plist.XMLFormat

type SequelAceRootPlist struct {
	FavoritesRoot struct {
		IsExpanded bool                 `plist:"IsExpanded"`
		Name       string               `plist:"Name"`
		Children   []SequelAcePlistItem `plist:"Children"`
	} `plist:"Favorites Root"`
}

type SequelAcePlistItem struct {
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

func NewSequelAceFlow(config *SequelAceDatasource) (spf *SequelAceFlow, err error) {
	bytes, err := ioutil.ReadFile(config.SequelAcePlistPath)
	if err != nil {
		return
	}

	sequelPlist := SequelAceRootPlist{}
	_, err = plist.Unmarshal(bytes, &sequelPlist)
	if err != nil {
		return
	}

	return &SequelAceFlow{
		SequelAceDatasource: *config,
		plist:               sequelPlist,
	}, nil
}

type SequelAceFlow struct {
	SequelAceDatasource
	plist       SequelAceRootPlist
	credentials *Credentials
}

func (sp *SequelAceFlow) findPlistItem() (plistItem *SequelAcePlistItem, foundItem bool) {
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

func (sp *SequelAceFlow) Name() string {
	return "SEQUEL ACE"
}

func (sp *SequelAceFlow) UpdateWithCredentials(credentials *Credentials) (err error) {
	plistItem, found := sp.findPlistItem()
	if !found {
		return errors.New(fmt.Sprintf("[SEQUEL ACE] Could not find plist item for database UUID %s", sp.DatabaseUUID))
	}

	plistItem.User = credentials.Username
	sp.credentials = credentials

	return nil
}

func (sp *SequelAceFlow) PersistChanges() (err error) {
	plistItem, found := sp.findPlistItem()
	if !found {
		return errors.New("Could not find a matching Sequel Database for that Sequel UUID.")
	}

	// These two values need to be synchronized for Sequel Ace to be able to
	// read the correct keychain value.
	service := fmt.Sprintf("Sequel Ace : %s (%d)", plistItem.Name, plistItem.Id)
	account := fmt.Sprintf("%s@%s/%s", plistItem.User, plistItem.Host, plistItem.Database)
	log.Printf("[SEQUEL ACE] Persisting username %s and password %s to service %s and account %s", plistItem.User, sp.credentials.Password, service, account)

	bytes, err := plist.Marshal(sp.plist, SEQUEL_ACE_PLIST_FORMAT)
	if err != nil {
		return
	}

	ioutil.WriteFile(sp.SequelAcePlistPath, bytes, 0644)

	return CreateOrUpdateKeychainEntriesForService(service, account, sp.credentials.Password, []string{})
}
