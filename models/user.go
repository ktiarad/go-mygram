package models

type User struct {
	GormModel
	Username string `gorm:"not null; unique; type:varchar(30)" json:"username" form:"username" valid:"required~Username is required"`
	Email    string `gorm:"not null; unique; type:varchar(30)" json:"email" form:"email" valid:"email~Valid email is required,required~Valid email is required"`
	Password string `gorm:"not null" json:"password" form:"password" valid:"required~Password is required,minstringlength(6)~Password has to have a minimum of 6 characters"`
	Age      int    `gorm:"not null" json:"age" form:"age" valid:"required~Age is required,min=9~Age should be above 8"`

	Photo       *[]Photo
	SocialMedia SocialMedia
	Comment     *[]Comment
}
