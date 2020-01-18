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


func executeAmanar() {
	configFilepath := os.Getenv("CONFIG_FILEPATH")
	if configFilepath == "" {
		log.Fatalln("[CONFIG FILEPATH] Please provide a configuration file as the environment variable CONFIG_FILEPATH so we can retrieve the Amanar configuration.")
	}
	configuration, err, resultErrors := amanar.LoadConfiguration(configFilepath)

	if err != nil {
		log.Fatalf("[CONFIG] Could not load configuration file: %s", err)
		return
	}

	if resultErrors != nil {
		amanar.HandleResultErrors(resultErrors)
	}

	githubToken := os.Getenv("GITHUB_TOKEN")
	if githubToken == "" {
		log.Fatalln("[GITHUB AUTH] Please provide a valid GitHub token as the environment variable GITHUB_TOKEN so we can fetch new credentials.")
		return
	}

	amanar.ProcessAmanar(githubToken, configuration)
}
