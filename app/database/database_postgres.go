package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"url_shortener/models"
)

func InitConnection() (*gorm.DB, error) {
	// dbURL should look like this
	// dbURL = "postgres://user:pass@host:db_port/db_name"
	dbURL := "postgres://"
	dbURL += os.Getenv("DB_USER")
	dbURL += ":" + os.Getenv("DB_PASS")
	dbURL += "@" + os.Getenv("DB_HOST")
	dbURL += ":" + os.Getenv("DB_PORT")
	dbURL += "/" + os.Getenv("DB_NAME")

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
		return db, err
	}

	db.AutoMigrate(&models.StorageURL{})

	return db, nil
}
