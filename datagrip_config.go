package main

import (
	"github.com/beevik/etree"
	"github.com/hashicorp/vault/api"
	"errors"
)

// Tested with DataGrip 2017.2
func NewDatagripConfig(filepath string) (*DatagripConfig, error) {
	d := etree.NewDocument()
	err := d.ReadFromFile(filepath)

	if err != nil {
		return nil, err
	}

	dc := &DatagripConfig{
		Document: d,
	}

	return dc, nil
}

// A DatagripConfig is an XML document containing
// This struct and methods allows updating of the username in
// such a configuration.
// DataGrip can store usernames in its configuration and passwords
// in the Keyring, or it can store both a username and password
// in a URL-like format in its config files. This updater assumes
// that the former is the case.
type DatagripConfig struct {
	Document *etree.Document
}

func (dc *DatagripConfig) UpdateUsername(databaseUuid string, secret *api.Secret) (oldUsername string, err error) {
	newUsername, ok := secret.Data["username"].(string)

	component := dc.Document.SelectElement("project").SelectElement("component")

	for _, dataSource := range component.SelectElements("data-source") {
		if uuid := dataSource.SelectAttrValue("uuid", ""); uuid == databaseUuid {

			if !ok {
				return "", errors.New("Could not update database UUID %s because secret lacked a username.")
			}

			username := dataSource.SelectElement("user-name")
			if username == nil {
				username = dataSource.CreateElement("user-name")
			}

			oldUsername = username.Text()
			username.SetText(newUsername)
			return oldUsername, nil

		}
	}

	return "", nil
}
