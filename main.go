package main

import (
	"log"
	"io/ioutil"
	"os"
	"encoding/json"
)

func main() {
	bytes, err := ioutil.ReadFile(os.Getenv("CONFIG_FILEPATH"))
	if err != nil {
		log.Fatal(err)
		return
	}

	configItems := []AmanarConfigItem{}
	err = json.Unmarshal(bytes, &configItems)
	if err != nil {
		log.Fatal(err)
		return
	}

	githubToken := os.Getenv("GITHUB_TOKEN")
	if githubToken == "" {
		log.Fatalln("Please provide a valid GitHub token as the environment variable GITHUB_TOKEN so we can fetch new credentials.")
		return
	}

	ghc := &VaultGithubAuthClient{
		GithubToken: githubToken,
	}
	err = ghc.loginWithGithub()
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, configItem := range configItems {
		secret, err := ghc.getCredential(configItem.VaultPath, configItem.VaultRole)
		if err != nil {
			log.Printf("Could not retrieve credential for vault path %s and vault role %s because %s. Skipping.", configItem.VaultPath, configItem.VaultRole, err)
		}
		ProcessConfigItem(&configItem.Configurables, secret)
	}
}
