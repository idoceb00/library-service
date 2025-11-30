package database

import (
	"log"

	"github.com/idoceb00/library-service/internal/models"
)

func Migrate() {
	err := DB.AutoMigrate(
		&models.Book{},
		&models.User{},
	)

	if err != nil {
		log.Fatalf("❌ Migration failed: %v", err)
	}

	log.Println("✅ Database migrated successfully")
}
