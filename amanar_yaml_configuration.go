package main

import "gopkg.in/yaml.v2"

func UnmarshalYamlAmanarConfiguration(data []byte) (AmanarConfiguration, error) {
	var r AmanarConfiguration
	err := yaml.Unmarshal(data, &r)
	return r, err
}
