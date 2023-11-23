package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Fullname  string    `gorm:"not null;type:varchar" json:"fullname"`
	Email     string    `gorm:"type:varchar;unique" json:"email"`
	Password  string    `gorm:"type:varchar" json:"password"`
	Role      string    `gorm:"type:varchar" json:"role"`
	Balance   int       `gorm:"not null;type:integer" json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserResponse struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Fullname  string    `gorm:"not null;type:varchar" json:"fullname"`
	Email     string    `gorm:"type:varchar;unique" json:"email"`
	Role      string    `gorm:"type:varchar" json:"role"`
	Balance   int       `gorm:"not null;type:integer" json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
