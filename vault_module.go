package vault_module

import (
	vault "github.com/hashicorp/vault/api"
)

func Login(addr string, token string) (*vault.Client, error){
	config := vault.DefaultConfig()

	config.Address = addr

	client, err := vault.NewClient(config)

	// Authenticate
	client.SetToken(token)

	return client, err
}




