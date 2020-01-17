package main

import (
	"fmt"
	"io"
	"text/template"
)


type TemplateSource struct {
	t *template.Template
	writer io.Writer
}

func (t TemplateSource) WriteToDisk(credentials Credentials) error {
	return t.t.Execute(t.writer, credentials)
}

func NewTemplateSourceFromFile(templateFilepath *string, writer io.Writer) (*TemplateSource, error) {
	t, err := template.ParseFiles(*templateFilepath)
	if err != nil {
		return nil, fmt.Errorf("could not parse template file: %w", err)
	}
	return &TemplateSource{
		t: t,
		writer: writer,
	}, nil
}


func NewTemplateSourceFromString(templateString *string, writer io.Writer) (*TemplateSource, error) {
	t, err := template.New("template_string").Parse(*templateString)
	if err != nil {
		return nil, fmt.Errorf("could not parse template string: %w", err)
	}

	return &TemplateSource{
		t: t,
		writer: writer,
	}, nil
}
