package main

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/vault/api"
)

func main() {
	config := api.DefaultConfig()
	config.Address = "http://127.0.0.1:8200"

	client, err := api.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to connect %s ", err)
	}

	// Authenticate
	client.SetToken("myroot")

	secretData := map[string]interface{}{
		"username": "root",
		"password": "root@123",
	}

	_, err = client.KVv2("secret").Put(context.Background(), "webapp", secretData)
	if err != nil {
		log.Fatalf("Failed to put %s ", err)
	}
	fmt.Println("Secret Written Successfully!")

	secret, err := client.KVv2("secret").Get(context.Background(), "webapp")
	if err != nil {
		log.Fatalf("Failed to get % s", err)
	}
	username, isUsernameOk := secret.Data["username"].(string)
	password, isPasswordOk := secret.Data["password"].(string)
	if !isUsernameOk || !isPasswordOk {
		log.Fatalf("Failed to get data")
	}
	fmt.Printf("Username %s \n", username)
	fmt.Printf("Password %s \n", password)
}
