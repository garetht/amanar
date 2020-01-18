package main

import (
	"github.com/garetht/amanar"
	"log"
	"os"
)

var GitCommit string
var BuildDate string

func main() {
	log.Printf("Amanar. version: %s, built: %s", GitCommit, BuildDate)
	executeAmanar()
}

func processAmanarConfigurationElement(githubToken string, ac amanar.AmanarConfiguration) {
	configurationProcessor, err := amanar.NewConfigurationProcessor(githubToken, ac)
	if err != nil {
		log.Fatalf("[PROCESS CONFIG] Could not process config: %s", err)
	}

	configurationProcessor.ProcessConfig()
}

func executeAmanar() {
	configurationElements, err, resultErrors := amanar.LoadConfiguration(os.Getenv("CONFIG_FILEPATH"))

	if err != nil {
		log.Fatalf("[CONFIG] Could not load configuration file: %s", err)
		return
	}

	if resultErrors != nil {
		log.Println("[CONFIG SCHEMA] The provided configuration did not conform to the structure required by the JSON Schema.")
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

	for _, configurationElement := range configurationElements {
		processAmanarConfigurationElement(githubToken, configurationElement)
	}
}
