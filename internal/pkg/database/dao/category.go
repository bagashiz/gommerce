package dao

import "gorm.io/gorm"

// Category is a struct that represent the category table in database
type Category struct {
	gorm.Model
	Name     string    `json:"name" gorm:"unique;not null"`
	Products []Product `json:"product"`
}
