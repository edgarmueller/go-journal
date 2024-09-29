package database

import (
	"log"

	"github.com/edgarmueller/go-api-journal/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var dbError error

func Connect(connectionString string) (db *gorm.DB) {
	Instance, dbError = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic(dbError)
	}
	log.Println("Connected to database")
	return Instance
}

func Migrate() {
	Instance.AutoMigrate(domain.User{}, domain.WorkJournal{}, domain.JournalEntry{})
	log.Println("Database migrated completed")
}
