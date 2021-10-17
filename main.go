package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	godotenv.Load(".env")
	apiKey := os.Getenv("API_KEY")
	apiSecret := os.Getenv("API_SECRET")

}
