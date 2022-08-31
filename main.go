package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	db := initPostgres()
	scrapeInfluenceId(db)
}

func initPostgres() *sql.DB {
	godotenv.Load()

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	database := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", user, password, host, database)

	// Connect to database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
