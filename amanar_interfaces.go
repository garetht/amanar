// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    amanar, err := UnmarshalAmanar(bytes)
//    bytes, err = amanar.Marshal()

package amanar

import (
	"fmt"
	"io"
)

type TemplateDefiner interface {
	GetTemplate() *string
	GetTemplatePath() *string
}

func NewTemplateSource(definer TemplateDefiner, writer io.Writer) (source *TemplateSource, err error) {
	if definer.GetTemplate() == nil && definer.GetTemplatePath() == nil {
		return nil, fmt.Errorf("neither template nor template path were provided")
	}

	if definer.GetTemplate() != nil && definer.GetTemplatePath() != nil {
		return nil, fmt.Errorf("provide at most one of template or template path")
	}

	if definer.GetTemplate() != nil {
		source, err = NewTemplateSourceFromString(definer.GetTemplate(), writer)
	} else {
		source, err = NewTemplateSourceFromFile(definer.GetTemplatePath(), writer)
	}

	if err != nil {
		return nil, fmt.Errorf("could not parse template: %w", err)
	}

	return
}

func (tds *TemplateDatasource) GetTemplate() *string {
	return tds.Template
}

func (tds *TemplateDatasource) GetTemplatePath() *string {
	return tds.TemplatePath
}

func (c *Constant) GetTemplate() *string {
	return c.Template
}

func (c *Constant) GetTemplatePath() *string {
	return c.TemplatePath
}
