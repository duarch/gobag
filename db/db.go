package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvVars() {
	err := godotenv.Load("credentials.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func ConnectDB() *sql.DB {
	GetEnvVars()
	dbName := os.Getenv("DB_NAME")
	dbSecret := os.Getenv("DB_KEY")
	openKey := "user=postgres password=" + dbSecret + " dbname=" + dbName + " host=localhost sslmode=disable"
	db, err := sql.Open("postgres", openKey)
	if err != nil {
		panic(err)
	}
	return db
}
