package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	ID             int    `gorm:"primaryKey;unique" json:"id"`
	Name           string `gorm:"not null" json:"name" form:"name" valid:"required~Name is required"`
	SocialMediaUrl string `gorm:"not null" json:"social_media_url" form:"social_media_url" valid:"required~Social media URL is required"`
	UserID         int    `json:"user_id"`
}

func (s *SocialMedia) BeforeCreate(tx *gorm.DB) error {
	_, errCreate := govalidator.ValidateStruct(s)

	if errCreate != nil {
		return errCreate
	}

	return nil
}
