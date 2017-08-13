package main

import (
	"log"
	"os"
)

func main() {
	configItems, err, resultErrors := LoadConfiguration(os.Getenv("CONFIG_FILEPATH"), "amanar_config_schema.json")

	if err != nil {
		log.Fatal(err)
		return
	}

	if resultErrors != nil {
		log.Println("[CONFIG SCHEMA] The provided configuration JSON did not conform to the structure required by the JSON Schema.")
		for _, resultErr := range resultErrors {
			log.Printf("[CONFIG SCHEMA] At JSON location %s: %s", resultErr.Context().String(), resultErr.Description())
		}
		return
	}

	githubToken := os.Getenv("GITHUB_TOKEN")
	if githubToken == "" {
		log.Fatalln("[GITHUB AUTH] Please provide a valid GitHub token as the environment variable GITHUB_TOKEN so we can fetch new credentials.")
		return
	}

	ghc := &VaultGithubAuthClient{
		GithubToken: githubToken,
	}
	err = ghc.loginWithGithub()
	if err != nil {
		log.Fatalf("[GITHUB AUTH] Could not log in with Github: %s", err)
		return
	}

	for _, configItem := range configItems {
		secret, err := ghc.getCredential(configItem.VaultPath, configItem.VaultRole)
		if err != nil {
			log.Printf("[VAULT AUTH] Could not retrieve credential for vault path %s and vault role %s because %s. Skipping.", configItem.VaultPath, configItem.VaultRole, err)
		}
		ProcessConfigItem(&configItem.Configurables, secret)
	}
}
