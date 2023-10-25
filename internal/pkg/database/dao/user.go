package dao

import (
	"time"
)

// User is a struct that represent the user table in database
type User struct {
	Model
	Name         string        `json:"name"`
	Password     string        `json:"password"`
	PhoneNumber  string        `json:"phone_number" gorm:"unique;not null"`
	Email        string        `json:"email" gorm:"unique;not null"`
	BirthDate    time.Time     `json:"birth_date" `
	About        string        `json:"about" gorm:"type:text"`
	Job          string        `json:"job"`
	ProvinceID   string        `json:"province_id"`
	CityID       string        `json:"city_id"`
	Addresses    []Address     `json:"address"`
	Shop         *Shop         `gorm:"foreignkey:UserID"`
	Transactions []Transaction `json:"transaction"`
	IsAdmin      bool          `json:"is_admin" gorm:"default:false"`
}
