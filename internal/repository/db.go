package repository

import (
	"log"
	"taskman/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func DBConnection(dsn string) (db *gorm.DB, err error) {
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error connectiong to database", err.Error())
	}

	if err = db.AutoMigrate(&models.Task{}); err != nil {
		log.Fatal("error migrating database")
	}

	return
}
