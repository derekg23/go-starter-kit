package models

import (
	"time"
)

type User struct {
	ID        uint64    `gorm:"primarykey;autoincrement"`
	Name      string    `gorm:"not null"`
	Email     string    `gorm:"not null;unique"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
