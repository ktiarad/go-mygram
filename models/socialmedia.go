package models

type SocialMedia struct {
	ID             int    `gorm:"primaryKey;unique" json:"id"`
	Name           string `gorm:"not null" json:"name" form:"name" valid:"required~Name is required"`
	SocialMediaUrl string `gorm:"not null" json:"social_media_url" form:"social_media_url" valid:"required~Social media URL is required"`
	UserID         int
}
