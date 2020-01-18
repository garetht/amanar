package amanar

import (
	"net/http"

	"errors"
	"fmt"

	"github.com/hashicorp/vault/api"
)


type GithubLoginBody struct {
	Token string `json:"token"`
}

type VaultGithubAuthClient struct {
	GithubToken string
	VaultAddress string
	vaultClient *api.Client
}

func (vc *VaultGithubAuthClient) LoginWithGithub() error {
	c, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		return err
	}
	err = c.SetAddress(vc.VaultAddress)
	if err != nil {
		return err
	}

	// The raw version requires the /v1, while the logical reads
	// do not need it.
	r := c.NewRequest(http.MethodPost, "/v1/auth/github/login")
	err = r.SetJSONBody(GithubLoginBody{Token: vc.GithubToken})
	if err != nil {
		return err
	}

	secret := api.Secret{}
	resp, err := c.RawRequest(r)
	if err != nil {
		return err
	}

	err = resp.DecodeJSON(&secret)
	if err != nil {
		return err
	}

	c.SetToken(secret.Auth.ClientToken)
	vc.vaultClient = c

	return nil
}

func (vc *VaultGithubAuthClient) GetCredential(vaultPath, vaultRole string) (*api.Secret, error) {
	if vc.vaultClient == nil || vc.vaultClient.Token() == "" {
		return nil, errors.New("Vault Github client has not yet been intialized with a token. Please log in.")
	}

	secret, err := vc.vaultClient.Logical().Read(fmt.Sprintf("%s/creds/%s", vaultPath, vaultRole))
	if err != nil {
		return nil, err
	}

	return secret, nil
}
