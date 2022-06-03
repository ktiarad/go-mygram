package services

import (
	"go-mygram/models"
	"go-mygram/params"
	"go-mygram/repositories"
	"net/http"
)

type UserServices struct {
	UserRepo repositories.UserRepo
}

func NewUserService(UserRepo repositories.UserRepo) *UserServices {
	return &UserServices{
		UserRepo: UserRepo,
	}
}

func (u *UserServices) Register(req *params.UserCreate) *params.Response {
	var user = &models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
		Age:      req.Age,
	}
	id, err := u.UserRepo.CreateUser(user)

	if err != nil {
		return &params.Response{
			Status:         http.StatusInternalServerError,
			Error:          "INTERNAL SERVER ERROR when creating user",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  http.StatusCreated,
		Message: "CREATED SUCCESS",
		Payload: id, // TODO : Payload berupa JSON berisi : age, email, id, username
	}
}

func (u *UserServices) Login(req *params.UserLogin) *params.Response {
	var user = &models.User{
		Email:    req.Email,
		Password: req.Password,
	}

	isOk, err := u.UserRepo.CheckUser(user)

	if err != nil {
		return &params.Response{
			Status:         http.StatusInternalServerError,
			Error:          "INTERNAL SERVER ERROR when login",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  http.StatusOK,
		Message: "LOGIN SUCCESS",
		Payload: isOk, // TODO : return payload token
	}

}

func (u *UserServices) UpdateUser(req *params.UserUpdate, id int) *params.Response {
	// TODO : headers Authorization (Bearer token string)
	// TODO : autentikasi dengan JWT

	var user = &models.User{
		Email:    req.Email,
		Username: req.Username,
	}

	err := u.UserRepo.UpdateUser(user, id)

	if err != nil {
		return &params.Response{
			Status:         http.StatusInternalServerError,
			Error:          "INTERNAL SERVER ERROR, when update user",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  http.StatusOK,
		Message: "UPDATE SUCCESS",
		Payload: id, // TODO : payload berupa id, email, username, age, updated_at
	}
}

func (u *UserServices) DeleteUser(id int) *params.Response {
	// TODO : header Authorization (Bearer token string)
	// TODO : autentikasi dengan JWT

	err := u.UserRepo.DeleteUser(id)

	if err != nil {
		return &params.Response{
			Status:         http.StatusInternalServerError,
			Error:          "INTERNAL SERVER ERROR, when delete user",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  http.StatusOK,
		Message: "Your account has been successfully deleted",
	}
}
