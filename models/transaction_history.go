package models

import "time"

type TransactionHistory struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	ProductID  uint      `gorm:"not null;type:integer" json:"product_id"`
	UserID     uint      `gorm:"not null;type:integer" json:"user_id"`
	Quantity   uint      `gorm:"not null;type:integer" json:"quantity"`
	TotalPrice int       `gorm:"not null;type:integer" json:"total_price"`
	Product    Product 	 `json:"product"`
	User       User 	 `json:"user"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}