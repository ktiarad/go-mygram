package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	ID        int `gorm:"primaryKey;unique" json:"id"`
	UserID    int
	PhotoID   int
	Message   string    `gorm:"not null" json:"message" form:"message" valid:"required~Message is required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (c *Comment) BeforeCreate(tx *gorm.DB) error {
	_, errCreate := govalidator.ValidateStruct(c)

	if errCreate != nil {
		return errCreate
	}

	return nil
}

// func (c *Comment) BeforeUpdate(tx *gorm.DB) error {
// 	_, errCreate := govalidator.ValidateStruct(c)

// 	if errCreate != nil {
// 		return errCreate
// 	}

// 	return nil
// }
