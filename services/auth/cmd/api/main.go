package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/felipemalacarne/lumina/auth/migrations"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const webPort = "8002"

var tries int64

type Config struct {
	DB *gorm.DB
}

func main() {
	log.Println("Starting server on port", webPort)

	db := connectToDB()
	if db == nil {
		log.Println("Failed to connect to database")
		return
	}

	migrations.DropTables(db)
	migrations.Migrate(db)

	// Ping the database to ensure the connection is alive
	pgDB, err := db.DB()
	if err != nil {
		log.Println("Failed to get database instance", err)
		return
	}

	err = pgDB.Ping()
	if err != nil {
		log.Println("Failed to ping database", err)
		return
	}

	app := Config{
		DB: db,
	}

	srv := &http.Server{
		Addr:    ":" + webPort,
		Handler: app.routes(),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Println("Server failed to start", err)
	}
}

func connectToDB() *gorm.DB {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")
	timezone := os.Getenv("DB_TIMEZONE")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		host, port, user, password, dbname, sslmode, timezone)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Failed to connect to database", err)
		return nil
	}

	return db
}
