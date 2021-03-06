package services

import (
	"go-mygram/helpers"
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

	payload := map[string]interface{}{
		"age":      req.Age,
		"email":    req.Email,
		"id":       id,
		"username": req.Username,
	}

	return &params.Response{
		Status:  http.StatusCreated,
		Message: "CREATED SUCCESS",
		Payload: payload,
	}
}

func (u *UserServices) Login(req *params.UserLogin) *params.Response {
	var user = &models.User{
		Email:    req.Email,
		Password: req.Password,
	}

	userDb, err := u.UserRepo.CheckUser(user) // Check data from database

	if err != nil {
		return &params.Response{
			Status:         http.StatusInternalServerError,
			Error:          "INTERNAL SERVER ERROR when login",
			AdditionalInfo: err.Error(),
		}
	}

	password := ""
	password = user.Password

	comparePass := helpers.ComparePass([]byte(userDb.Password), []byte(password))
	if !comparePass {
		return &params.Response{
			Status: http.StatusUnauthorized,
			Error:  "Invalid email/password",
		}
	}

	token, err := helpers.GenerateToken(userDb.Email, userDb.ID)

	if err != nil {
		return &params.Response{
			Status: http.StatusInternalServerError,
			Error:  "INTERNAL SERVER ERROR, when generate token",
		}
	}

	payload := map[string]interface{}{
		"token": token,
	}

	return &params.Response{
		Status:  http.StatusOK,
		Message: "LOGIN SUCCESS",
		Payload: payload,
	}

}

func (u *UserServices) UpdateUser(req *params.UserUpdate, id int) *params.Response {

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

	userDb, err := u.UserRepo.GetUserById(id)

	if err != nil {
		return &params.Response{
			Status:         http.StatusNotFound,
			Error:          "NOT FOUND",
			AdditionalInfo: err.Error(),
		}
	}

	payload := map[string]interface{}{
		"id":         id,
		"email":      userDb.Email,
		"username":   userDb.Username,
		"age":        userDb.Age,
		"updated_at": userDb.UpdatedAt,
	}

	return &params.Response{
		Status:  http.StatusOK,
		Message: "UPDATE SUCCESS",
		Payload: payload,
	}
}

func (u *UserServices) DeleteUser(id int) *params.Response {
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
