package controllers

import (
	"go-mygram/params"
	"go-mygram/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *services.UserServices
}

func NewUserController(userService *services.UserServices) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (u *UserController) Register(ctx *gin.Context) {
	var req params.UserCreate

	err := ctx.BindJSON(&req)

	if err != nil {
		response := params.Response{
			Status:         http.StatusBadRequest,
			Error:          "BAD REQUEST, when register",
			AdditionalInfo: err.Error(),
		}

		params.WriteJsonResponse(ctx.Writer, &response)
	}

	response := u.userService.Register(&req)

	params.WriteJsonResponse(ctx.Writer, response)
}

func (u *UserController) Login(ctx *gin.Context) {
	var req params.UserLogin

	err := ctx.BindJSON(&req)

	if err != nil {
		response := params.Response{
			Status:         http.StatusBadRequest,
			Error:          "BAD REQUEST, when login",
			AdditionalInfo: err.Error(),
		}

		params.WriteJsonResponse(ctx.Writer, &response)
	}

	response := u.userService.Login(&req)

	params.WriteJsonResponse(ctx.Writer, response)
}

func (u *UserController) UpdateUser(ctx *gin.Context) {
	var req params.UserUpdate

	userID := ctx.Param("userID")
	id, err := strconv.Atoi(userID)

	if err != nil {
		response := params.Response{
			Status:         http.StatusBadRequest,
			Error:          "BAD REQUEST, when get param user id",
			AdditionalInfo: err.Error(),
		}
		params.WriteJsonResponse(ctx.Writer, &response)
	}

	err = ctx.BindJSON(&req)

	if err != nil {
		response := params.Response{
			Status:         http.StatusBadRequest,
			Error:          "BAD REQUEST, when update user",
			AdditionalInfo: err.Error(),
		}

		params.WriteJsonResponse(ctx.Writer, &response)
	}

	response := u.userService.UpdateUser(&req, id)
	params.WriteJsonResponse(ctx.Writer, response)
}

func (u *UserController) DeleteUser(ctx *gin.Context) {
	userID := ctx.Param("userID")
	id, err := strconv.Atoi(userID)

	if err != nil {
		response := params.Response{
			Status:         http.StatusBadRequest,
			Error:          "BAD REQUEST, when delete user",
			AdditionalInfo: err.Error(),
		}
		params.WriteJsonResponse(ctx.Writer, &response)
	}

	response := u.userService.DeleteUser(id)
	params.WriteJsonResponse(ctx.Writer, response)
}
