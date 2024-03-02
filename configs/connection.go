package connection

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func ConnectDB() {
	dbUser := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	var err error
	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatalf("Terjadi kesalahan koneksi database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Server database tidak merespon: %v", err)
	}

	fmt.Println("Berhasil terhubung ke database!")
}

func GetDB() *sql.DB {
	return db
}
