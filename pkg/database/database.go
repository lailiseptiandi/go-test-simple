package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/lailiseptiandi/go-test-simple/internal/models"
	"github.com/lailiseptiandi/go-test-simple/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB(c *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", c.DB_HOST, c.DB_USERNAME, c.DB_PASSWORD, c.DB_NAME, c.DB_PORT, c.TimeZone)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			Colorful:                  true,
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}
	MigrateDB(db)

	AddColumn(db, &models.Payment{}, "status")

	return db, nil
}

func DisconnectDB(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to kill connection from database")
	}
	dbSQL.Close()
}

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&models.Payment{})
}

func AddColumn(db *gorm.DB, dst interface{}, column string) {
	db.Migrator().AddColumn(dst, column)
}
