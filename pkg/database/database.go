package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/mongj/gds-onecv-swe-assignment/ent"
	"github.com/mongj/gds-onecv-swe-assignment/ent/enttest"
)

var Client *ent.Client

// Init initializes the database connection and schema
func Init() {
	var err error
	Client, err = ent.Open(
		os.Getenv("DB_DIALECT"),
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PASSWORD")))

	if err != nil {
		log.Fatalf("Failed to establish connection to database: %v", err)
	} else {
		log.Println("Established connection to database")
	}

	// Run the auto migration tool.
	if err := Client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("Failed to create schema resources: %v", err)
	} else {
		log.Println("Created schema resource")
	}
}

// InitTest initializes an in-memory sqlite database for testing
func InitTest(t *testing.T) {
	Client = enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
}
