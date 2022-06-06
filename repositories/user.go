package repositories

import (
	"go-mygram/models"

	"gorm.io/gorm"
)

type UserRepo interface {
	CreateUser(request *models.User) (int, error)
	CheckUser(request *models.User) (*models.User, error)
	UpdateUser(request *models.User, id int) error
	GetUserById(id int) error
	DeleteUser(id int) error
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (u *userRepo) CreateUser(request *models.User) (int, error) {
	result := u.db.Create(request)
	err := result.Error
	id := request.ID

	return id, err
}

func (u *userRepo) CheckUser(request *models.User) (*models.User, error) {
	var user models.User

	result := u.db.Where("email=?", request.Email).First(&user)
	err := result.Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userRepo) UpdateUser(request *models.User, id int) error {
	var user models.User

	result := u.db.Model(&user).Where("id=?", id).Updates(models.User{
		Email:    request.Email,
		Username: request.Username,
	})
	err := result.Error

	return err
}

func (u *userRepo) GetUserById(id int) error {
	var user models.User

	result := u.db.First("id=?", id, &user)

	err := result.Error

	return err
}

func (u *userRepo) DeleteUser(id int) error {
	// TODO : autentikasi dengan JWT
	var user models.User

	result := u.db.Model(&user).Where("id=?", id).Delete(&user)
	err := result.Error

	return err
}
