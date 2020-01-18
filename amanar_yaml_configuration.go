package amanar

import "gopkg.in/yaml.v2"

func UnmarshalYamlAmanarConfiguration(data []byte) (Amanar, error) {
	var r Amanar
	err := yaml.Unmarshal(data, &r)
	return r, err
}

func (r *AmanarConfiguration) MarshalYaml() ([]byte, error) {
	return yaml.Marshal(r)
}
