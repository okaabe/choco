package adapter

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectAdapter(dialector gorm.Dialector) (*gorm.DB, error) {
	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second, // Slow SQL threshold
				LogLevel:                  logger.Info, // Log level
				IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
				Colorful:                  true,        // Disable color
			},
		),
	})

	if err != nil {
		return nil, err
	}

	return db, nil
}

func ConnectProductionAdapter(dsn string) (*gorm.DB, error) {
	return ConnectAdapter(postgres.Open(dsn))
}

func ConnectTestAdapter() (*gorm.DB, error) {
	return ConnectAdapter(sqlite.Open(":memory:"))
}
