package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
