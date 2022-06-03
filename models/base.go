package models

import "time"

type GormModel struct {
	ID        int       `gorm:"primaryKey;unique" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TODO : buat Hooks
