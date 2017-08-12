package main

import (
	"io/ioutil"
	"log"
	"fmt"
	"encoding/json"
)

func main() {
	if err != nil {
		log.Fatal(err)
	}
	config := []AmanarConfigItem{}
	json.Unmarshal(bytes, &config)
	fmt.Printf("%#v", config)

	/*	ghc := &VaultGithubAuthClient{
			GithubToken: os.Getenv("GITHUB_TOKEN"),
		}
		err := ghc.loginWithGithub()
		if err != nil {
			log.Fatal(err)
			return
		}

		"./example/backup.xml"
		filename := "/Users/garethtan/Library/Preferences/DataGrip2017.2/projects/default/.idea/dataSources.local.xml"

		source := IntellijDatasource{
			DatabaseUUID:       "1f76a8f4-2328-4ee3-ad77-d95ba60c645b",
			DatasourceFilepath: filename,
			VaultPath:          "db-tzanalytic",
			VaultRole:          "read-only",
			AuthedClient:       ghc,
		}

		err = source.UpdateCredentials()
		if err != nil {
			log.Fatal(err)
			return
		}

		source = IntellijDatasource{
			DatabaseUUID:       "73d00a82-3e4d-4978-ae34-2bb04755429c",
			DatasourceFilepath: filename,
			VaultPath:          "db-peakpass",
			VaultRole:          "read-only",
			AuthedClient:       ghc,
		}

		err = source.UpdateCredentials()
		if err != nil {
			log.Fatal(err)
			return
		}*/
}
