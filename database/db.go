package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	// ConnectDB connects to the database

	connStr := os.Getenv("DB")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	runMigrations(db)
	return db
}

func runMigrations(db *sql.DB) {
	// Read SQL commands from the migrations file
	content, err := os.ReadFile("migrations/createtables.sql")
	if err != nil {
		log.Fatalf("Failed to read the migration file: %v", err)
	}

	// Split SQL commands
	commands := string(content)

	// Execute each command
	_, err = db.Exec(commands)
	if err != nil {
		log.Fatalf("Failed to execute the migration: %v", err)
	}

	fmt.Println("Migration run successfully!")
}
