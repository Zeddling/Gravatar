package main

import (
	"context"
	"log"
	"os"

	"github.com/blockfrost/blockfrost-go"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	log.Println("Project Id: ", os.Getenv("PROJECT_ID"))
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{
			ProjectID: os.Getenv("PROJECT_ID"),
			Server:    os.Getenv("CARDANO_TESTNET"),
		},
	)
	info, err := client.Info(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("API Info: ", info.Url)
	log.Println("API Version: ", info.Version)
}
