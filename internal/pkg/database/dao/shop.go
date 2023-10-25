package dao

// Shop is a struct that represent the shop table in databasek
type Shop struct {
	Model
	Name           string    `json:"name"`
	ProfilePicture string    `json:"profile_picture"`
	UserID         uint      `json:"user_id" gorm:"not null"`
	Products       []Product `json:"product"`
}
