# Vault Client module for Golang

This provides the [git.assistanz.com/kishore/vault_module](https://git.assistanz.com/kishore/vault_module) package which contains some useful functions for interacting with a Vault server.

## Basic usage

```
    package main

    import (
        "fmt"
        "log"
        "git.assistanz.com/kishore/vault_module"
    )

    func main() {
        config := vault.DefaultConfig()
        config.Address = "http://<vault_addr:8200/"

        // Authenticate
	    client.SetToken("<Auth Token>")

        // Create new secret in vault
        err = secrets.(client, "test", "new")
        if (err != nil) {
		    fmt.Print(err.Error())
	    }
    }
```
