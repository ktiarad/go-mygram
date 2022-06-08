package controllers

import (
	"go-mygram/params"
	"go-mygram/services"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type SocialMediaController struct {
	socialMediaServices *services.SocialMediaServices
}

func NewSocialMediaController(socialMediaServices *services.SocialMediaServices) *SocialMediaController {
	return &SocialMediaController{
		socialMediaServices: socialMediaServices,
	}
}

func (s *SocialMediaController) CreateSocialMedia(ctx *gin.Context) {
	var socialmedia params.SocialMediaCreate

	err := ctx.BindJSON(&socialmedia)

	if err != nil {
		response := params.Response{
			Status:         http.StatusBadRequest,
			Error:          "BAD REQUEST, when create social media",
			AdditionalInfo: err.Error(),
		}
		params.WriteJsonResponse(ctx.Writer, &response)
	}

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := int(userData["userID"].(float64))

	socialmedia.UserID = userID

	response := s.socialMediaServices.CreateSocialMedia(&socialmedia)
	params.WriteJsonResponse(ctx.Writer, response)
}

func (s *SocialMediaController) GetAllSocialMedias(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := int(userData["userID"].(float64))

	response := s.socialMediaServices.GetAllSocialMedias(userID)
	params.WriteJsonResponse(ctx.Writer, response)
}

func (s *SocialMediaController) UpdateSocialMedia(ctx *gin.Context) {
	var socialmedia params.SocialMediaCreate

	err := ctx.BindJSON(&socialmedia)

	if err != nil {
		response := &params.Response{
			Status:         http.StatusBadRequest,
			Error:          "BAD REQUEST, when update social media",
			AdditionalInfo: err.Error(),
		}

		params.WriteJsonResponse(ctx.Writer, response)
	}

	socialmediaId := ctx.Param("socialMediaId")
	id, err := strconv.Atoi(socialmediaId)

	response := s.socialMediaServices.UpdateSocialMedia(&socialmedia, id)
	params.WriteJsonResponse(ctx.Writer, response)
}

func (s *SocialMediaController) DeleteSocialMedia(ctx *gin.Context) {
	socialmediaID := ctx.Param("socialMediaId")
	id, err := strconv.Atoi(socialmediaID)

	if err != nil {
		response := &params.Response{
			Status:         http.StatusBadRequest,
			Error:          "BAD REQUEST, when get param social media id",
			AdditionalInfo: err.Error(),
		}

		params.WriteJsonResponse(ctx.Writer, response)
	}

	response := s.socialMediaServices.DeleteSocialMedia(id)
	params.WriteJsonResponse(ctx.Writer, response)
}
