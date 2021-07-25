package core

import (
	models "basic/core/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(cfg postgres.Config) error {
	db, err := gorm.Open(
		postgres.New(cfg),
	)
	if err == nil {
		DB = db
	}
	return err
}

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&models.User{})
	return err
}
