package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/metgag/itrix-challenge/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	host := os.Getenv("PG_HOST")
	user := os.Getenv("PG_USER")
	pwd := os.Getenv("PG_PASSWORD")
	dbname := os.Getenv("PG_DB")
	port := os.Getenv("PG_PORT")
	timezone := os.Getenv("PG_TZ")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", host, user, pwd, dbname, port, timezone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("unable to connect to dabatase: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("unable to get sql.db: %v", err)
	}
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err := db.AutoMigrate(
		&model.Todo{},
		&model.Category{},
	); err != nil {
		return nil, fmt.Errorf("failed to run migrations: %v", err)
	}

	log.Println("DB connected and migrated succesfully")
	return db, nil
}
