package main

import (
	connection "crud-golang-mysql/configs"
	"crud-golang-mysql/routes"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Terjadi kesalahan saat memuat file .env: %v", err)
	}

	connection.ConnectDB()

	router := routes.SetupRoutes()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server berjalan:", "http://localhost:"+port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
