package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	ID        int       `gorm:"primaryKey;unique" json:"id"`
	Title     string    `gorm:"not null; type:varchar(30)" json:"title" form:"title" valid:"required~Title is required"`
	Caption   string    `json:"caption" form:"caption"`
	PhotoUrl  string    `gorm:"not null" json:"photo_url" form:"photo_url" valid:"required~Photo URL is required"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Comment *[]Comment
}

func (p *Photo) BeforeCreate(tx *gorm.DB) error {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		return errCreate
	}

	return nil
}
