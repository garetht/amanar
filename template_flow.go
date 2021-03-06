package amanar

import (
	"fmt"
	"io"
	"log"
)

func NewTemplateFlow(config *TemplateDatasource, writer io.Writer) (*TemplateFlow, error) {
	source, err := NewTemplateSource(config, writer)
	if err != nil {
		return nil, fmt.Errorf("could not create template flow from template datasource: %s", err)
	}

	return &TemplateFlow{
		source:      source,
	}, nil
}

type TemplateFlow struct {
	credentials *Credentials
	source      *TemplateSource
}

func (tf *TemplateFlow) Name() string {
	return "TEMPLATE FLOW"
}

func (tf *TemplateFlow) UpdateWithCredentials(credentials *Credentials) error {
	log.Printf("[%s DATASOURCE] Updating template flow with new username %s and password %s", tf.Name(), credentials.Username, credentials.Password)
	tf.credentials = credentials
	log.Printf("[%s DATASOURCE] Updated template flow with new username %s and password %s", tf.Name(), credentials.Username, credentials.Password)
	return nil
}

func (tf *TemplateFlow) PersistChanges() error {
	log.Printf("[%s DATASOURCE] Writing new username %s and password %s to template in stdout", tf.Name(), tf.credentials.Username, tf.credentials.Password)
	// prints to the configured Writer
	if err := tf.source.WriteToDisk(*tf.credentials); err != nil {
		return err
	}
	log.Printf("[%s DATASOURCE] Successfully wrote new username %s and password %s to template in stdout", tf.Name(), tf.credentials.Username, tf.credentials.Password)
	return nil
}
