package main

import (
	"github.com/beevik/etree"
	"fmt"
)

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

			for _, attr := range dataSource.Attr {
				fmt.Printf("  ATTR: %s=%s\n", attr.Key, attr.Value)
			}
		}
	}
}
