package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func CreateUser(db *gorm.DB, user *User) (*User, error) {
	result := db.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func GetUserByUsername(db *gorm.DB, username string) (*User, error) {
	var user *User
	result := db.Where(&User{Username: username}).First(&user)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func GetUserById(db *gorm.DB, id uuid.UUID) (*User, error) {
	var user *User
	result := db.Where(&User{ID: id}).First(&user)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
