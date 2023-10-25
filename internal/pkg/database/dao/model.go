package dao

import (
	"time"

	"gorm.io/gorm"
)

// Model is a basic physical data model for data objects.
type Model struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
