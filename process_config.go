package amanar

import (
	"fmt"
	"io"
	"log"
)

type ConfigurationProcessor interface {
	ProcessConfig()
}

type VaultConfigurationProcessor struct {
	credentials        *Credentials
	vaultClient        VaultClient
	vaultAddress       string
	vaultConfiguration []VaultConfiguration
	writer             io.Writer
}

func (v VaultConfigurationProcessor) ProcessConfig() {
	log.Printf("\n\n\n\n =========================== [VAULT ADDRESS %s] =========================== \n\n", v.vaultAddress)

	err := v.vaultClient.Login()
	if err != nil {
		log.Fatalf("[VAULT AUTH] Could not log in to Vault: %s", err)
		return
	}

	for _, configItem := range v.vaultConfiguration {
		secret, err := v.vaultClient.GetCredential(configItem.VaultPath, configItem.VaultRole)
		if err != nil {
			log.Printf("[VAULT AUTH] Could not retrieve secret for vault path %s and vault role %s because %s. Skipping.", configItem.VaultPath, configItem.VaultRole, err)
			continue
		}

		credentials, err := CreateCredentialsFromSecret(secret)

		if err != nil {
			log.Printf("[VAULT AUTH] Could not convert Vault secret into Amanar credentials because %s. Skipping.", err)
			continue
		}

		log.Printf("[VAULT CONFIGURATION] %v:%v", configItem.VaultPath, configItem.VaultRole)
		ProcessVaultConfigItem(&configItem.Configurables, credentials, v.writer)
	}
}

type ConstantConfigurationProcessor struct {
	constant Constant
	writer   io.Writer
}

func (c ConstantConfigurationProcessor) ProcessConfig() {
	ProcessConstantConfigItem(c.constant, c.writer)
}

func NewConfigurationProcessor(githubToken string, ac AmanarConfiguration, writer io.Writer) (ConfigurationProcessor, error) {
	if ac.Constant == nil && ac.VaultAddress == nil && ac.VaultConfiguration == nil {
		return nil, fmt.Errorf("please provide either a Constant configuration or a Vault configuration")
	}

	if ac.Constant != nil && ac.VaultAddress != nil && ac.VaultConfiguration != nil {
		return nil, fmt.Errorf("please provide only one of a Constant configuration and a Vault configuration")
	}

	if ac.VaultAddress != nil && ac.VaultConfiguration != nil {
		var vc VaultClient = nil

		// for backward compatibility, vault_auth is optional in the config file
		if ac.VaultAuth == nil {
			if githubToken == "" {
				vc = &DefaultVaultAuthClient{
					VaultAddress: *ac.VaultAddress,
				}
			} else {
				vc = &VaultGithubAuthClient{
					GithubToken:  githubToken,
					VaultAddress: *ac.VaultAddress,
				}
			}
		} else {
			switch *ac.VaultAuth {
			case "token":
				vc = &DefaultVaultAuthClient{
					VaultAddress: *ac.VaultAddress,
				}
			case "github":
				vc = &VaultGithubAuthClient{
					GithubToken:  githubToken,
					VaultAddress: *ac.VaultAddress,
				}
			case "aws_iam":
				vaultDBUsernamePrefix := ""
				if ac.VaultDBUsernamePrefix != nil {
					vaultDBUsernamePrefix = *ac.VaultDBUsernamePrefix
				}
				vc = &VaultAwsIamAuthClient{
					VaultAddress:          *ac.VaultAddress,
					VaultDBUsernamePrefix: vaultDBUsernamePrefix,
				}
			default:
				return nil, fmt.Errorf("unsupported Vault auth method %s", *ac.VaultAuth)
			}
		}

		return VaultConfigurationProcessor{
			vaultClient:        vc,
			vaultAddress:       *ac.VaultAddress,
			vaultConfiguration: ac.VaultConfiguration,
			writer:             writer,
		}, nil
	}

	if ac.Constant != nil {
		return ConstantConfigurationProcessor{
			constant: *ac.Constant,
			writer:   writer,
		}, nil
	}

	return nil, fmt.Errorf("please provide a full Constant configuration or a full Vault configuration")
}

func ProcessConstantConfigItem(constant Constant, writer io.Writer) {
	var errs []error

	templateSource, err := NewTemplateSource(&constant, writer)
	if err != nil {
		errs = append(errs, fmt.Errorf("could not create new template source: %w", err))
	}

	err = templateSource.WriteToDiskWithoutContext()
	if err != nil {
		errs = append(errs, fmt.Errorf("could not write constant to disk: %w", err))
	}

	if len(errs) > 0 {
		for _, err := range errs {
			log.Printf("[CONSTANT PROCESSING] Encountered errors processing constant: %#v. Processing constants that worked.", err)
		}
	}
}

func ProcessVaultConfigItem(configurables *Configurables, credentials *Credentials, writer io.Writer) {
	var errs []error
	var flows []Flower

	for _, datasourceConfig := range configurables.IntellijDatasources {
		flow, err := NewIntellijDatasourceFlow(&datasourceConfig)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		flows = append(flows, flow)
	}

	for _, runConfigurationsConfig := range configurables.IntellijRunConfigurations {
		flow, err := NewIntellijRunConfigsFlow(&runConfigurationsConfig)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		flows = append(flows, flow)
	}

	for _, querious2Config := range configurables.Querious2Datasources {
		flow, err := NewQuerious2Flow(&querious2Config)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		flows = append(flows, flow)
	}

	for _, sequelProConfig := range configurables.SequelProDatasources {
		flow, err := NewSequelProFlow(&sequelProConfig)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		flows = append(flows, flow)
	}

	for _, sequelAceConfig := range configurables.SequelAceDatasources {
		flow, err := NewSequelAceFlow(&sequelAceConfig)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		flows = append(flows, flow)
	}

	for _, posticoConfig := range configurables.PosticoDatasources {
		flow, err := NewPosticoFlow(&posticoConfig)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		flows = append(flows, flow)
	}

	for _, shellConfig := range configurables.ShellDatasources {
		flow, err := NewShellFlow(&shellConfig)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		flows = append(flows, flow)
	}

	for _, jsonConfig := range configurables.JSONDatasources {
		flow, err := NewJSONFlow(&jsonConfig)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		flows = append(flows, flow)
	}

	for _, templateConfig := range configurables.TemplateDatasources {
		flow, err := NewTemplateFlow(&templateConfig, writer)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		flows = append(flows, flow)
	}

	if len(errs) > 0 {
		for _, err := range errs {
			log.Printf("[FLOW PROCESSING] Encountered errors processing flow: %#v. Processing flows that worked.", err)
		}
	}

	UpdateCredentials(flows, credentials)

	return
}

func UpdateCredentials(flows []Flower, credentials *Credentials) {
	var err error
	for _, flow := range flows {
		log.Printf("[%s] Beginning to update flow %#v with credentials %s", flow.Name(), flow, credentials)

		err = flow.UpdateWithCredentials(credentials)
		if err != nil {
			log.Printf("[%s] Error when performing non-write update to flow %#v with credentials %s. Will not try and persist externally. Skipping ahead to next flow.", flow.Name(), flow, credentials)
			log.Print(err)
			continue
		}

		err = flow.PersistChanges()
		if err != nil {
			log.Printf("[%s] Error when persisting changes to to flow %#v with credentials %s. Skipping ahead to next flow.", flow.Name(), flow, credentials)
			log.Print(err)
		}
	}
}
