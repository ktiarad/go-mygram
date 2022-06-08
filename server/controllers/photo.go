package controllers

import (
	"go-mygram/params"
	"go-mygram/services"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type PhotoController struct {
	photoService *services.PhotoServices
}

func NewPhotoController(photoService *services.PhotoServices) *PhotoController {
	return &PhotoController{
		photoService: photoService,
	}
}

func (p *PhotoController) CreatePhoto(ctx *gin.Context) {
	var req params.PhotoCreate

	err := ctx.BindJSON(&req)

	if err != nil {
		response := params.Response{
			Status:         http.StatusBadRequest,
			Error:          "BAD REQUEST, when create photo",
			AdditionalInfo: err.Error(),
		}

		params.WriteJsonResponse(ctx.Writer, &response)
	}

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := int(userData["userID"].(float64))

	req.UserID = userID

	response := p.photoService.CreatePhoto(&req)

	params.WriteJsonResponse(ctx.Writer, response)
}

func (p *PhotoController) GetAllPhotos(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := int(userData["userID"].(float64))

	response := p.photoService.GetAllPhotos(userID)

	params.WriteJsonResponse(ctx.Writer, response)
}

func (p *PhotoController) UpdatePhoto(ctx *gin.Context) {
	var req params.PhotoCreate

	photoId := ctx.Param("photoID")
	id, err := strconv.Atoi(photoId)

	if err != nil {
		response := params.Response{
			Status:         http.StatusBadRequest,
			Error:          "BAD REQUEST, when get param photo id",
			AdditionalInfo: err.Error(),
		}

		params.WriteJsonResponse(ctx.Writer, &response)
	}

	err = ctx.BindJSON(&req)

	if err != nil {
		response := params.Response{
			Status:         http.StatusBadRequest,
			Error:          "BAD REQUEST, when update photo",
			AdditionalInfo: err.Error(),
		}
		params.WriteJsonResponse(ctx.Writer, &response)
	}

	response := p.photoService.UpdatePhoto(&req, id)
	params.WriteJsonResponse(ctx.Writer, response)
}

func (p *PhotoController) DeletePhoto(ctx *gin.Context) {
	photoId := ctx.Param("photoID")
	id, err := strconv.Atoi(photoId)

	if err != nil {
		response := params.Response{
			Status: http.StatusBadRequest,
			Error:  "BAD REQUEST, when get param photo id",
		}
		params.WriteJsonResponse(ctx.Writer, &response)
	}

	response := p.photoService.DeletePhoto(id)
	params.WriteJsonResponse(ctx.Writer, response)

}
