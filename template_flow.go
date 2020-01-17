package amanar

import (
	"fmt"
	"log"
	"os"
)

func NewTemplateFlow(config *TemplateDatasource) (*TemplateFlow, error) {
	if config.Template == nil && config.TemplatePath == nil {
		return nil, fmt.Errorf("neither template nor template path were provided")
	}

	if config.Template != nil && config.TemplatePath != nil {
		return nil, fmt.Errorf("provide at most one of template or template path")
	}


	var err error
	var source *TemplateSource

	writer := os.Stdout
	if config.Template != nil {
		source, err = NewTemplateSourceFromString(config.Template, writer)
	} else {
		source, err = NewTemplateSourceFromFile(config.TemplatePath, writer)
	}

	if err != nil {
		return nil, fmt.Errorf("couldn not parse template: %w", err)
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
	// print to stdout
	log.Printf("[%s DATASOURCE] Writing new username %s and password %s to template in stdout", tf.Name(), tf.credentials.Username, tf.credentials.Password)
	if err := tf.source.WriteToDisk(*tf.credentials); err != nil {
		return err
	}
	log.Printf("[%s DATASOURCE] Successfully wrote new username %s and password %s to template in stdout", tf.Name(), tf.credentials.Username, tf.credentials.Password)
	return nil
}
