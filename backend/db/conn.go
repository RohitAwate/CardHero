package db

import (
	"CardHero/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var database *gorm.DB

func init() {
	host := os.Getenv("CH_PG_HOST")
	port := os.Getenv("CH_PG_PORT")
	user := os.Getenv("CH_PG_USER")
	pass := os.Getenv("CH_PG_PASS")
	dbName := os.Getenv("CH_PG_DBNAME")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s TimeZone=America/New_York",
		host, user, pass, dbName, port,
	)

	var err error
	database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	log.Println("Connected to database")

	log.Println("Running auto-migrations")
	err = database.AutoMigrate(models.GetAll()...)
	if err != nil {
		panic(err)
	}

	log.Println("Setting up triggers")
	models.SetupTriggers(database)

	log.Println("Database initialized")
}

func getConn() *gorm.DB {
	return database
}
