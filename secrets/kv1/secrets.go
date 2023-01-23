package kv1

import (
	"context"
	"fmt"
	"log"

	vault "github.com/hashicorp/vault/api"
)

func Get(client *vault.Client, mountPath string, key string ) ( *vault.KVSecret, error) {

	secret, err := client.KVv1(mountPath).Get(context.Background(), key)
	if err != nil {
		log.Fatalf("unable to write secret: %v", err)
	}

	return secret, err
}

func Create(client *vault.Client, mountPath string, key string , value string )  error {
	secretData := map[string]interface{}{
		key: value,
	}
	err := client.KVv1(mountPath).Put(context.Background(), key, secretData)
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Printf("written successfully")
	return  nil
}

func Update(client *vault.Client, key string , value string ) error {
	secretData := map[string]interface{}{
		key: value,
	}
	err := client.KVv1("secret/rundeck/keys").Put(context.Background(), key, secretData)
	if err != nil {
		log.Fatalf("unable to write secret: %v", err)
	}
	fmt.Printf("written successfully")
	return nil
}


