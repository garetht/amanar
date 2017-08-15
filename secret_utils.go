package main

import (
	"github.com/hashicorp/vault/api"
	"errors"
)

type Credentials struct {
	Username string
	Password string
}

func CreateCredentialsFromSecret(secret *api.Secret) (*Credentials, error) {
	newUsername, oku := secret.Data["username"].(string)
	newPassword, okp := secret.Data["password"].(string)

	if !oku {
		return nil, errors.New("[CREDENTIAL] Could not parse username out from Vault secret.")
	}

	if !okp {
		return nil, errors.New("[CREDENTIAL] Could not parse password out from Vault secret.")
	}

	return &Credentials{
		Username: newUsername,
		Password: newPassword,
	}, nil
}
