package amanar

import (
	"net/http"

	"errors"
	"fmt"

	"github.com/hashicorp/vault/api"

	"io/ioutil"
	"os/user"
)

type VaultClient interface {
	Login() error
	GetCredential(vaultPath, vaultRole string) (*api.Secret, error)
}

type GithubLoginBody struct {
	Token string `json:"token"`
}

type VaultGithubAuthClient struct {
	GithubToken  string
	VaultAddress string
	vaultClient  *api.Client
}

type DefaultVaultAuthClient struct {
	VaultAddress string
	vaultClient  *api.Client
}

func (dc *DefaultVaultAuthClient) Login() error {
	c, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		return err
	}
	err = c.SetAddress(dc.VaultAddress)
	if err != nil {
		return err
	}

	usr, err := user.Current()
	if err != nil {
		return err
	}

	blob, err := ioutil.ReadFile(fmt.Sprintf("%s/.vault-token", usr.HomeDir))
	if err != nil {
		return err
	}
	token := string(blob)
	fmt.Printf("Read Vault token: %s\n", token)

	c.SetToken(token)
	dc.vaultClient = c

	return nil
}

func (vc *VaultGithubAuthClient) Login() error {
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

func (vc *DefaultVaultAuthClient) GetCredential(vaultPath, vaultRole string) (*api.Secret, error) {
	if vc.vaultClient == nil || vc.vaultClient.Token() == "" {
		return nil, errors.New("Vault client has not yet been intialized with a token. Please log in.")
	}

	secret, err := vc.vaultClient.Logical().Read(fmt.Sprintf("%s/creds/%s", vaultPath, vaultRole))
	if err != nil {
		return nil, err
	}

	return secret, nil
}
func (vc *VaultGithubAuthClient) GetCredential(vaultPath, vaultRole string) (*api.Secret, error) {
	if vc.vaultClient == nil || vc.vaultClient.Token() == "" {
		return nil, errors.New("Vault client has not yet been intialized with a token. Please log in.")
	}

	secret, err := vc.vaultClient.Logical().Read(fmt.Sprintf("%s/creds/%s", vaultPath, vaultRole))
	if err != nil {
		return nil, err
	}

	return secret, nil
}
