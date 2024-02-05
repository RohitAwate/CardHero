package db

import (
	"CardHero/models"
	"CardHero/monitoring"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

	var monitor monitoring.Monitor = monitoring.NewPrintMonitor("db/conn.go#init()")
	monitor.LogInfo("Connected to database")

	monitor.LogInfo("Running auto-migrations")
	err = database.AutoMigrate(models.GetAll()...)
	if err != nil {
		panic(err)
	}

	monitor.LogInfo("Setting up triggers")
	models.SetupTriggers(database)

	monitor.LogInfo("Database initialized")
}

func getConn() *gorm.DB {
	return database
}
