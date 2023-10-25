package dao

// Address is a struct that represent the address table in database
type Address struct {
	Model
	Title       string `json:"title"`
	Receiver    string `json:"receiver"`
	PhoneNumber string `json:"phone_number"`
	Details     string `json:"details" gorm:"type:text"`
	UserID      uint   `json:"user_id" gorm:"not null"`
}
