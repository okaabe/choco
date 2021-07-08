package adapters

import (
	"choco/server/internals/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB(dialector gorm.Dialector) *gorm.DB {
	db, err := gorm.Open(dialector, &gorm.Config{})

	if err != nil {
		log.Fatalf("Couldn't connect to the database cause: %s", err)
	}

	db.AutoMigrate(&models.Community{}, &models.User{}, &models.Post{})

	return db
}

func ConnectTestDB() *gorm.DB {
	return ConnectDB(sqlite.Open(":memory:"))
}

func ConnectProdDB() *gorm.DB {
	return ConnectDB(postgres.Open(os.Getenv("POSTGRES_DSN")))
}
