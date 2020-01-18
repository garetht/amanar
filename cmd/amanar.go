package main

import (
	"github.com/garetht/amanar"
	"log"
	"os"
)

var GitCommit string
var GitTag string
var BuildDate string

func main() {
	log.Printf("Amanar %s. sha: %s. built: %s.\n\n", GitTag, GitCommit, BuildDate)
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
		return
	}

	// The Github token is only needed if the configuration specifies
	// Vault credentials are needed, which we will check for later.
	githubToken := os.Getenv("GITHUB_TOKEN")
	amanar.ProcessAmanar(githubToken, configuration)
}
