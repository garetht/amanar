package amanar

import (
	"github.com/beevik/etree"
	"net/url"
	"log"
)

func NewIntellijRunConfig(filepath string) (*IntellijRunConfig, error) {
	d := etree.NewDocument()
	err := d.ReadFromFile(string(filepath))

	if err != nil {
		return nil, err
	}

	rc := &IntellijRunConfig{
		Document: d,
		Fullpath: filepath,
	}

	return rc, nil
}

type IntellijRunConfig struct {
	Document *etree.Document
	Fullpath string
}

func (rc *IntellijRunConfig) UpdateEnvironmentVariable(environmentVariable, databaseHost string, credentials *Credentials) (err error) {
	envVars := rc.Document.SelectElement("component").SelectElement("configuration").SelectElement("envs")
	for _, envVar := range envVars.SelectElements("env") {
		if envVarName := envVar.SelectAttrValue("name", ""); envVarName == environmentVariable {

			value := envVar.SelectAttrValue("value", "")
			updatedValue, err := createOrUpdateUsernamePasswordWithHost(value, databaseHost, credentials.Username, credentials.Password)

			if err != nil {
				return err
			}
			envVar.CreateAttr("value", updatedValue)
		}
	}

	return nil
}

func (rc *IntellijRunConfig) WriteToFile() (err error) {
	log.Printf("[RUN CONFIGS] Writing new run configuration to file %s", rc.Fullpath)
	return rc.Document.WriteToFile(rc.Fullpath)
}

func createOrUpdateUsernamePasswordWithHost(urlValue, databaseHost, username, password string) (string, error) {
	parsedUrl, err := url.Parse(urlValue)
	if err != nil {
		return "", err
	}

	if databaseHost == parsedUrl.Host {
		log.Printf("[RUN CONFIGS] Updating URL %s with username %s and password %s", urlValue, username, password)

		newUserInfo := url.UserPassword(username, password)
		parsedUrl.User = newUserInfo
	}

	return parsedUrl.String(), nil
}
