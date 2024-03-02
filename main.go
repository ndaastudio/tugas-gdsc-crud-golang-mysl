package main

import (
	connection "crud-golang-mysql/configs"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Terjadi kesalahan saat memuat file .env: %v", err)
	}

	connection.ConnectDB()
}
