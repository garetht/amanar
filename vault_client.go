package amanar

import (
	"net/http"

	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/hashicorp/vault/api"

	"io/ioutil"
	"os/user"

	"encoding/base64"
	"encoding/json"
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

type VaultAwsIamAuthClient struct {
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
	if vc.GithubToken == "" {
		return fmt.Errorf("[GITHUB AUTH] Please provide a valid GitHub token as the environment variable GITHUB_TOKEN so we can fetch new credentials")
	}

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

func (vc *VaultAwsIamAuthClient) Login() error {
	c, err := api.NewClient(&api.Config{Address: vc.VaultAddress})
	if err != nil {
		return err
	}

	loginData := make(map[string]interface{})

	stsSession := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	var params *sts.GetCallerIdentityInput
	svc := sts.New(stsSession)
	stsRequest, _ := svc.GetCallerIdentityRequest(params)

	stsRequest.Sign()

	headersJson, err := json.Marshal(stsRequest.HTTPRequest.Header)
	if err != nil {
		return err
	}
	requestBody, err := ioutil.ReadAll(stsRequest.HTTPRequest.Body)
	if err != nil {
		return err
	}

	loginData["iam_http_request_method"] = stsRequest.HTTPRequest.Method
	loginData["iam_request_url"] = base64.StdEncoding.EncodeToString([]byte(stsRequest.HTTPRequest.URL.String()))
	loginData["iam_request_headers"] = base64.StdEncoding.EncodeToString(headersJson)
	loginData["iam_request_body"] = base64.StdEncoding.EncodeToString(requestBody)

	secret, err := c.Logical().Write("auth/aws/login", loginData)
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

func (vc *VaultAwsIamAuthClient) GetCredential(vaultPath, vaultRole string) (*api.Secret, error) {
	if vc.vaultClient == nil || vc.vaultClient.Token() == "" {
		return nil, errors.New("Vault client has not yet been intialized with a token. Please log in.")
	}

	secret, err := vc.vaultClient.Logical().Read(fmt.Sprintf("%s/creds/%s", vaultPath, vaultRole))
	if err != nil {
		return nil, err
	}

	return secret, nil
}
