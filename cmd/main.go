package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/metgag/itrix-challenge/internal/config"
	"github.com/metgag/itrix-challenge/internal/router"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := config.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	router := router.InitRoute(db)
	router.Run(":6080")
}
