package cmd

import (
	"context"
	"errors"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/security/keyvault/azsecrets"
)

type azureKeyVault struct {
	client *azsecrets.Client
	cache  map[string]string
}

func (v *azureKeyVault) URL(vaultName string) string {
	return fmt.Sprintf("https://%s.vault.azure.net/", vaultName)
}

type azureKeyVaultConfig struct {
	DefaultVault string `json:"defaultVault" mapstructure:"defaultVault" yaml:"defaultVault"`
	vaults       map[string]*azureKeyVault
	cred         *azidentity.DefaultAzureCredential
}

func (a *azureKeyVaultConfig) GetSecret(secretName, vaultName string) string {
	if a.vaults == nil {
		a.vaults = make(map[string]*azureKeyVault)
	}

	if _, ok := a.vaults[vaultName]; !ok {
		a.vaults[vaultName] = &azureKeyVault{}
	}

	if secret, ok := a.vaults[vaultName].cache[secretName]; ok {
		return secret
	}

	if a.cred == nil {
		a.cred = mustValue(azidentity.NewDefaultAzureCredential(nil))
	}

	if a.vaults[vaultName].client == nil {
		a.vaults[vaultName].client = mustValue(azsecrets.NewClient(a.vaults[vaultName].URL(vaultName), a.cred, nil))
	}

	resp := mustValue(a.vaults[vaultName].client.GetSecret(context.Background(), secretName, "", nil))

	if a.vaults[vaultName].cache == nil {
		a.vaults[vaultName].cache = make(map[string]string)
	}

	a.vaults[vaultName].cache[secretName] = *resp.Value

	return *resp.Value
}

func (c *Config) azureKeyVaultTemplateFunc(args ...string) string {
	var secretName, vaultName string

	switch len(args) {
	case 1:
		if c.AzureKeyVault.DefaultVault == "" {
			panic(errors.New("no value set in azureKeyVault.defaultVault"))
		}
		secretName, vaultName = args[0], c.AzureKeyVault.DefaultVault
	case 2:
		secretName, vaultName = args[0], args[1]
	default:
		panic(fmt.Errorf("expected 1 or 2 arguments, got %d", len(args)))
	}

	return c.AzureKeyVault.GetSecret(secretName, vaultName)
}
