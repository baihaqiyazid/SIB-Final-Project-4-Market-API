package models

import "time"

type Product struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Title      string    `gorm:"not null;type:varchar" json:"title"`
	Price      int       `gorm:"not null;type:integer" json:"price"`
	Stock      int      `gorm:"not null;type:integer" json:"stock"`
	CategoryID uint      `gorm:"not null;type:integer" json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
