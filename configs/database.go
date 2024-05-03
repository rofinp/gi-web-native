package configs

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

const (
	DB_USER = "developer"
	DB_PASS = "supersecretpassword"
	DB_NAME = "go_products"
	DB_HOST = "localhost"
	DB_PORT = "5432"
)

func ConnectToDatabase() {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta",
		DB_HOST, DB_PORT, DB_USER, DB_PASS, DB_NAME,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	log.Println("Connection established on database")
	DB = db
}
