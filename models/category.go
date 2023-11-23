package models

import "time"

type Category struct {
	ID                uint      `gorm:"primaryKey" json:"id"`
	Type              string    `gorm:"not null;type:varchar" json:"type"`
	SoldProductAmount int       `gorm:"not null;type:integer" json:"sold_product_amount"`
	Products          []Product `json:"products"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
