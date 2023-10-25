package dao

// Product is a struct that represent the product table in database
type Product struct {
	Model
	Name          string         `json:"name"`
	Slug          string         `json:"slug"`
	ResellerPrice int            `json:"reseller_price"`
	RetailPrice   int            `json:"retail_price"`
	Stock         int            `json:"stock"`
	Description   string         `json:"description" gorm:"type:text"`
	ShopID        uint           `json:"shop_id" gorm:"not null"`
	Shop          *Shop          `json:"shop"`
	CategoryID    uint           `json:"category_id" gorm:"not null"`
	Category      *Category      `json:"category"`
	Photos        []ProductPhoto `json:"photos" gorm:"foreignKey:ProductID"`
}

// ProductLog is a struct that represent the product_logs table in database
type ProductLog struct {
	Model
	ProductID     uint           `json:"product_id" gorm:"not null"`
	Name          string         `json:"name"`
	Slug          string         `json:"slug"`
	ResellerPrice int            `json:"reseller_price"`
	RetailPrice   int            `json:"retail_price"`
	Description   string         `json:"description" gorm:"type:text"`
	ShopID        uint           `json:"shop_id" gorm:"not null"`
	Shop          *Shop          `json:"shop"`
	CategoryID    uint           `json:"category_id" gorm:"not null"`
	Category      *Category      `json:"category"`
	Photos        []ProductPhoto `json:"photos" gorm:"foreignKey:ProductLogID"`
}

// ProductPhoto is a struct that represent the product_photos table in database
type ProductPhoto struct {
	Model
	Url          string `json:"url"`
	ProductID    uint   `json:"product_id" gorm:"not null;index"`
	ProductLogID uint   `json:"product_log_id" gorm:"not null;index"`
}
