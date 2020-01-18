package amanar

import (
	"github.com/xeipuuv/gojsonschema"
	"io"
	"log"
	"os"
)

func ProcessAmanarWithWriter(githubToken string, a *Amanar, writer io.Writer) {
	for _, configurationElement := range a.AmanarConfiguration {
		configurationProcessor, err := NewConfigurationProcessor(githubToken, configurationElement, writer)
		if err != nil {
			log.Fatalf("[PROCESS CONFIG] Could not process config: %s", err)
		}

		configurationProcessor.ProcessConfig()
	}
}

// This is the main entrypoint for Amanar. Given a Github Token and
// a valid Amanar configuration, performs the necessary side effects.
func ProcessAmanar(githubToken string, a *Amanar) {
	ProcessAmanarWithWriter(githubToken, a, os.Stdout)
}

func HandleResultErrors(resultErrors []gojsonschema.ResultError) {
		log.Println("[CONFIG SCHEMA VALIDATION] The provided configuration did not conform to the structure required by the JSON Schema.")
		for _, resultErr := range resultErrors {
			log.Printf("[CONFIG SCHEMA VALIDATION] At JSON location %s: %s", resultErr.Context().String(), resultErr.Description())
		}
}
