package dao

// Category is a struct that represent the category table in database
type Category struct {
	Model
	Name     string    `json:"name" gorm:"unique;not null"`
	Products []Product `json:"product"`
}
