package main

import "github.com/beevik/etree"

// Tested with DataGrip 2017.2
func NewDatagripConfig(document []byte) (*DatagripConfig, error) {
	d := etree.NewDocument()
	err := d.ReadFromBytes(document)

	if err != nil {
		return nil, err
	}

	dc := &DatagripConfig{
		Document: d,
	}

	return dc, nil
}

type DatagripConfig struct {
	Document *etree.Document
}

func (dc *DatagripConfig) UpdateUsername(databaseUuid, newUsername string) {
	component := dc.Document.SelectElement("project").SelectElement("component")

	for _, dataSource := range component.SelectElements("data-source") {
		if uuid := dataSource.SelectAttrValue("uuid", ""); uuid == databaseUuid {

			username := dataSource.SelectElement("user-name")
			if username != nil {
				username.SetText(newUsername)
			}
		}
	}
}
