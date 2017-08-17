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


type IntellijRunConfigsFlow struct {
	IntellijRunConfigurationsConfig
	credentials *Credentials
	runConfigurations []*IntellijRunConfig
}

func (rc *IntellijRunConfigsFlow) UpdateWithCredentials(credentials *Credentials) (err error) {
	err = rc.parseRunConfigs()
	if err != nil {
		return
	}

	rc.credentials = credentials

	for _, runConfig := range rc.runConfigurations {
		runConfig.UpdateEnvironmentVariable(rc.EnvironmentVariable, rc.DatabaseHost, credentials)
	}

	return nil
}

func (rc *IntellijRunConfigsFlow) parseRunConfigs() (err error) {
	if rc.runConfigurations != nil {
		return nil
	}

	files, err := ioutil.ReadDir(rc.RunConfigurationsFolderPath)

	if err != nil {
		return
	}

	rcs := []*IntellijRunConfig{}
	for _, file := range files {
		if name := file.Name(); filepath.Ext(name) == ".xml" && !file.IsDir() {
			fullPath := filepath.Join(rc.RunConfigurationsFolderPath, name)
			rc, err := NewIntellijRunConfig(fullPath)
			if err != nil {
				log.Printf("[RUN CONFIGS] Could not parse run config %s. Skipping.", fullPath)
			} else {
				rcs = append(rcs, rc)
			}
		}
	}


	rc.runConfigurations = rcs

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

func (rc *IntellijRunConfigsFlow) UpdateCredentials(credentials *Credentials) (err error) {
	err = rc.UpdateCredentials(credentials)
	if err != nil {
		return
	}

	err = rc.PersistChanges()
	if err != nil {
		return
	}

	return nil
}
