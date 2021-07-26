package services

import (
	models "basic/core/models"
	"fmt"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var once sync.Once

func InitDB(cfg postgres.Config) error {
	var err error
	once.Do(func() {
		_db, err := gorm.Open(
			postgres.New(cfg),
		)
		if err == nil && db == nil {
			db = _db
		}
	})

	return err
}

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&models.User{})
	return err
}

func GetDB() (*gorm.DB, error) {
	if db == nil {
		return nil, fmt.Errorf("DB not instantiated")
	}
	return db, nil
}
