package models

import "time"

type Comment struct {
	ID        int `gorm:"primaryKey;unique" json:"id"`
	UserID    int
	PhotoID   int
	Message   string    `gorm:"not null" json:"message" form:"message" valid:"required~Message is required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TODO : buat Hooks BeforeCreate
