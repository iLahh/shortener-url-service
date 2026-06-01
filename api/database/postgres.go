package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectPostgres() {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PG_HOST"),
		os.Getenv("PG_PORT"),
		os.Getenv("PG_USER"),
		os.Getenv("PG_PASS"),
		os.Getenv("PG_NAME"),
	)

	var err error
	DB, err = sql.Open("postgres", dsn)

	if err != nil {
		log.Fatal("Failed to open postgres connection:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Failed to ping postgres:", err)
	}

	fmt.Println("PostgreSQL connected")
	runMigrations()
}

func runMigrations() {
	query := `
	CREATE TABLE IF NOT EXISTS urls (
        id          SERIAL PRIMARY KEY,
        short_code  VARCHAR(50) UNIQUE NOT NULL,
        long_url    TEXT NOT NULL,
        expiry      BIGINT NOT NULL DEFAULT 24,
        click_count BIGINT NOT NULL DEFAULT 0,
        created_at  TIMESTAMP DEFAULT NOW()
    );`

	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal("Migration failed: ", err)
	}
	log.Println("Migration completed")
}
