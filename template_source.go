package main

import (
	"fmt"
	"os"
	"text/template"
)


type TemplateSource struct {
	t *template.Template
}

func (t TemplateSource) WriteToDisk(credentials Credentials) error {
	return t.t.Execute(os.Stdout, credentials)
}

func NewTemplateSourceFromFile(templateFilepath *string) (*TemplateSource, error) {
	t, err := template.ParseFiles(*templateFilepath)
	if err != nil {
		return nil, fmt.Errorf("could not parse template file: %w", err)
	}
	return &TemplateSource{
		t: t,
	}, nil
}


func NewTemplateSourceFromString(templateString *string) (*TemplateSource, error) {
	t, err := template.New("template_string").Parse(*templateString)
	if err != nil {
		return nil, fmt.Errorf("could not parse template string: %w", err)
	}

	return &TemplateSource{
		t: t,
	}, nil
}
