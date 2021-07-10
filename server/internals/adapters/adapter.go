package adapters

import (
	"choco/server/internals/models"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB(dialector gorm.Dialector) *gorm.DB {
	logger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:              time.Second,   // Slow SQL threshold
			LogLevel:                   logger.Info, // Log level
			IgnoreRecordNotFoundError: true,           // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,          // Disable color
		},
	)

	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: logger,
	})

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
