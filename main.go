package main

import (
	"fmt"
	"log"
	"vault_client/secrets/kv1"

	vault "github.com/hashicorp/vault/api"
)

func main() {

	config := vault.DefaultConfig()

	config.Address = "http://23.82.14.238:8200/"

	client, err := vault.NewClient(config)
	if err != nil {
		log.Fatalf("unable to initialize Vault client: %v", err)
	}

	// Authenticate
	client.SetToken("hvs.CAESIATYKZ6p8w5tbAvtgF8Gy8yPfe2bmUDdV3lERcQesoGkGh4KHGh2cy52TmI4c3pXTjZJOFdubDVTSkVpNmhpWlU")
	
	err = secrets.(client, "test", "new")
	if (err != nil) {
		fmt.Print(err.Error())
	}

	// secrets.Create(client, "test" , "test")
}
