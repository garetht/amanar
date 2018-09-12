package main

import (
	"io/ioutil"
	"log"
	"path/filepath"
)

// UpdateUsername - must not have side effects
// UpdatePassword - must not have side effects
// In the database these actions are simply queued for later execution
// Persist


func NewIntellijRunConfigsFlow(config *IntellijRunConfiguration) (*IntellijRunConfigsFlow, error) {
	files, err := ioutil.ReadDir(config.RunConfigurationsFolderPath)

	if err != nil {
		return nil, err
	}

	rcs := []*IntellijRunConfig{}
	for _, file := range files {
		if name := file.Name(); filepath.Ext(name) == ".xml" && !file.IsDir() {
			fullPath := filepath.Join(config.RunConfigurationsFolderPath, name)
			rc, err := NewIntellijRunConfig(fullPath)
			if err != nil {
				log.Printf("[RUN CONFIGS] Could not parse run config %s. Skipping.", fullPath)
			} else {
				rcs = append(rcs, rc)
			}
		}
	}

	return &IntellijRunConfigsFlow{
		IntellijRunConfiguration: *config,
		runConfigurations:         rcs,
	}, nil
}

type IntellijRunConfigsFlow struct {
	IntellijRunConfiguration
	credentials *Credentials
	runConfigurations []*IntellijRunConfig
}

func (rc *IntellijRunConfigsFlow) Name() string {
	return "INTELLIJ RUN CONFIGURATION"
}

func (rc *IntellijRunConfigsFlow) UpdateWithCredentials(credentials *Credentials) (err error) {
	rc.credentials = credentials

	for _, runConfig := range rc.runConfigurations {
		runConfig.UpdateEnvironmentVariable(rc.EnvironmentVariable, rc.DatabaseHost, credentials)
	}

	return nil
}

func (rc *IntellijRunConfigsFlow) PersistChanges() (err error) {
	if err != nil {
		return
	}

	for _, runConfig := range rc.runConfigurations {
		err = runConfig.WriteToFile()
		if err != nil {
			log.Printf("[RUN CONFIGS] Error writing %s to file. Skipping.", runConfig.Fullpath)
		}
	}

	return nil
}
