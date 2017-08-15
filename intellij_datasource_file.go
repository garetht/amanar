package main

import (
	"github.com/beevik/etree"
)

// Tested with DataGrip 2017.2
func NewIntellijDatasourceFile(filepath IntellijDatasourceFilepath) (*IntellijDatasourceFile, error) {
	d := etree.NewDocument()
	err := d.ReadFromFile(string(filepath))

	if err != nil {
		return nil, err
	}

	dc := &IntellijDatasourceFile{
		Document: d,
		Fullpath: filepath,
	}

	return dc, nil
}

// A IntellijDatasourceFile is an XML document containing
// This struct and methods allows updating of the username in
// such a configuration.
// DataGrip can store usernames in its configuration and passwords
// in the Keyring, or it can store both a username and password
// in a URL-like format in its config files. This updater assumes
// that the former is the case.
type IntellijDatasourceFile struct {
	Document *etree.Document
	Fullpath IntellijDatasourceFilepath
}

func (dc *IntellijDatasourceFile) UpdateUsername(databaseUuid IntellijDatabaseUUID, credentials *Credentials) (oldUsername string, err error) {
	component := dc.Document.SelectElement("project").SelectElement("component")

	for _, dataSource := range component.SelectElements("data-source") {
		if uuid := dataSource.SelectAttrValue("uuid", ""); IntellijDatabaseUUID(uuid) == databaseUuid {

			username := dataSource.SelectElement("user-name")
			if username == nil {
				username = dataSource.CreateElement("user-name")
			}

			oldUsername = username.Text()
			username.SetText(credentials.Username)
			return oldUsername, nil
		}
	}

	return "", nil
}

func (dc *IntellijDatasourceFile) WriteToFile() error {
	return dc.Document.WriteToFile(string(dc.Fullpath))
}
