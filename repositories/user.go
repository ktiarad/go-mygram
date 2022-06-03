package repositories

import (
	"go-mygram/models"

	"gorm.io/gorm"
)

type UserRepo interface {
	CreateUser(request *models.User) (int, error)
	CheckUser(request *models.User) (bool, error)
	UpdateUser(request *models.User, id int) error
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
	// TODO : hashing password dengan bycript
	// TODO : buat Hook
	result := u.db.Create(request)
	err := result.Error
	id := request.ID

	return id, err
}

func (u *userRepo) CheckUser(request *models.User) (bool, error) {
	// TODO : cek password dengan bycript
	var user models.User

	result := u.db.Where("email=?", request.Email).Where("password", request.Password).Find(&user)
	err := result.Error

	if err != nil {
		return false, err
	}

	if &user == nil {
		return false, err
	}

	return true, nil
}

func (u *userRepo) UpdateUser(request *models.User, id int) error {
	// TODO : autentikasi dengan JWT
	var user models.User

	result := u.db.Model(&user).Where("id=?", id).Updates(models.User{
		Email:    request.Email,
		Username: request.Username,
	})
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
