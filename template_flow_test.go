package main

import (
	"bytes"
	"os"
	"testing"
)

func TestTemplateFlow_PersistChanges(t *testing.T) {
	type fields struct {
		credentials *Credentials
		source      *TemplateSource
	}
	strBuf := bytes.NewBuffer([]byte{})
	s := "Hello from {{.Username}} and {{.Password}}"
	templSource, _ := NewTemplateSourceFromString(&s, strBuf)

	fileBuf := bytes.NewBuffer([]byte{})
	wd, _ := os.Getwd()
	path := wd + "/example/example_template_file.go.md"
	fileSource, err := NewTemplateSourceFromFile(&path, fileBuf)
	if err != nil {
		print("hello")
	}
	tests := []struct {
		name       string
		buffer     *bytes.Buffer
		fields     fields
		wantErr    error
		wantOutput string
	}{
		{name: "Can write a string template to a writer",
			buffer: strBuf,
			fields: struct {
				credentials *Credentials
				source      *TemplateSource
			}{credentials: &Credentials{Username: "uname", Password: "password"}, source: templSource},
			wantErr:    nil,
			wantOutput: "Hello from uname and password"},
		{name: "Can write a file template to stdout",
			buffer: fileBuf,
			fields: struct {
				credentials *Credentials
				source      *TemplateSource
			}{credentials: &Credentials{Username: "firstname", Password: "secondname"}, source: fileSource},
			wantErr: nil,
			wantOutput: `Template
===

The credentials are firstname and secondname
`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tf := &TemplateFlow{
				credentials: tt.fields.credentials,
				source:      tt.fields.source,
			}
			if err := tf.PersistChanges(); err != tt.wantErr {
				t.Errorf("PersistChanges() error = %v, wantErr %v", err, tt.wantErr)
			}

			output := tt.buffer.String()
			if output != tt.wantOutput {
				t.Errorf("output error, received output = %v, want output = %v", output, tt.wantOutput)
			}
		})
	}
}
