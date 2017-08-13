package main

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/hashicorp/vault/api"
)

type IntellijRunConfigsFlow struct {
	IntellijRunConfigurationsConfig
	NewVaultSecret *api.Secret
}

func (rc *IntellijRunConfigsFlow) parseRunConfigs() (rcs []*IntellijRunConfig, err error) {
	files, err := ioutil.ReadDir(rc.RunConfigurationsFolderPath)

	if err != nil {
		return
	}

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

	return
}

func (rc *IntellijRunConfigsFlow) UpdateCredentials() (err error) {
	runConfigs, err := rc.parseRunConfigs()

	if err != nil {
		return
	}

	for _, runConfig := range runConfigs {
		runConfig.UpdateEnvironmentVariable(rc.EnvironmentVariable, rc.DatabaseHost, rc.NewVaultSecret)
	}

	for _, runConfig := range runConfigs {
		err = runConfig.WriteToFile()
		if err != nil {
			log.Printf("[RUN CONFIGS] Error writing %s to file. Skipping.", runConfig.Fullpath)
		}
	}

	return nil
}
