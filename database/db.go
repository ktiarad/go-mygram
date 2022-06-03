package database

import (
	"fmt"
	"go-mygram/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	DB_HOST = "localhost"
	DB_PORT = "5432"
	DB_USER = "ktiarad"
	DB_PASS = "123456"
	DB_NAME = "mygram"
)

func ConnectDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASS, DB_NAME)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	log.Default().Println("Connection db success")

	err = migration(db)
	if err != nil {
		panic(err)
	}

	return db
}

func migration(db *gorm.DB) error {
	if err := db.AutoMigrate(models.User{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(models.Photo{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(models.Comment{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(models.SocialMedia{}); err != nil {
		return err
	}

	return nil
}
