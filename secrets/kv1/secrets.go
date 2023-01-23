package kv1

import (
	"context"

	vault "github.com/hashicorp/vault/api"
)

func Get(client *vault.Client, mountPath string, key string ) ( *vault.KVSecret, error) {

	secret, err := client.KVv1(mountPath).Get(context.Background(), key)
	return secret, err
}

func Create(client *vault.Client, mountPath string, key string , value string )  error {
	secretData := map[string]interface{}{
		key: value,
	}
	err := client.KVv1(mountPath).Put(context.Background(), key, secretData)
	return  err
}

func Update(client *vault.Client,mountPath string, key string , value string ) error {
	secretData := map[string]interface{}{
		key: value,
	}
	err := client.KVv1(mountPath).Put(context.Background(), key, secretData)
	return  err
}

func Delete(client *vault.Client, mountPath string, key string ) error {
	err := client.KVv1(mountPath).Delete(context.Background(), key)
	return  err
}


