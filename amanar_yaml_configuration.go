package amanar

import "gopkg.in/yaml.v2"


//go:generate bash -c "npx quicktype --src-lang schema amanar_config_schema.json --package amanar --top-level Amanar --lang go | sed -E -e 's/json:\"(.+)\"/json:\"\\1\" yaml:\"\\1\"/g' > amanar_configuration.go"
func UnmarshalYamlAmanarConfiguration(data []byte) (Amanar, error) {
	var r Amanar
	err := yaml.Unmarshal(data, &r)
	return r, err
}

func DynamicUnmarshalYamlAmanarConfiguration(data []byte) (map[interface{}]interface{}, error) {
	var r map[interface{}]interface{}
	err := yaml.Unmarshal(data, &r)
	return r, err
}

func (r *AmanarConfiguration) MarshalYaml() ([]byte, error) {
	return yaml.Marshal(r)
}
