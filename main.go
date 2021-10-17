package main

import (
	"ftx-crypto/ftx"
	"github.com/joho/godotenv"
	"os"
)

func getFtxClient() *ftx.FtxClient {
	apiKey := os.Getenv("API_KEY")
	apiSecret := os.Getenv("API_SECRET")
	client := ftx.New(apiKey, apiSecret, "")

	return client
}

func main() {
	godotenv.Load(".env")
	//
	//ftxClient := getFtxClient()
	//c := ftxClient
}
