package models

import (
	"go-mygram/helpers"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	ID        int       `gorm:"primaryKey;unique" json:"id"`
	Username  string    `gorm:"not null; unique; type:varchar(30)" json:"username" form:"username" valid:"required~Username is required"`
	Email     string    `gorm:"not null; unique; type:varchar(30)" json:"email" form:"email" valid:"email~Valid email is required,required~Valid email is required"`
	Password  string    `gorm:"not null" json:"password" form:"password" valid:"required~Password is required,minstringlength(6)~Password has to have a minimum of 6 characters"`
	Age       int       `gorm:"not null" json:"age" form:"age" valid:"required~Age is required,range(9|100)~Age has to be above 8"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Photo       *[]Photo
	SocialMedia *SocialMedia
	Comment     *[]Comment
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		return errCreate
	}

	u.Password = helpers.HashPass(u.Password)

	return nil
}

// func (u *User) BeforeUpdate(tx *gorm.DB) error {
// 	_, errCreate := govalidator.ValidateStruct(u)

// 	if errCreate != nil {
// 		return errCreate
// 	}

// 	return nil
// }
