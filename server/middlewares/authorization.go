package middlewares

import (
	"go-mygram/database"
	"go-mygram/models"
	"go-mygram/repositories"
	"log"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func UserAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := database.ConnectDB()
		userRepo := repositories.NewUserRepo(db)

		userIdParam, err := strconv.Atoi(ctx.Param("userID"))

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":           "BAD REQUEST, when authorizing",
				"additional_info": err.Error(),
			})
		}

		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userId := int(userData["id"].(int))

		_, err = userRepo.GetUserById(userId)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":           "NOT FOUND",
				"additional_info": err.Error(),
			})
		}

		log.Default().Println("id user : ", userId)

		if userIdParam != userId {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "UNAUTHORIZED, when authorizing",
			})
		}

		ctx.Next()

	}
}

func PhotoAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := database.ConnectDB()
		photoRepo := repositories.NewPhotoRepo(db)

		photoIdParam, err := strconv.Atoi(ctx.Param("photoID"))

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":           "BAD REQUEST, when authorizing",
				"additional_info": err.Error(),
			})
		}

		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userId := int(userData["userID"].(float64))
		log.Default().Println("userID:", userId)

		var photo *models.Photo
		photo, err = photoRepo.GetPhotoById(photoIdParam)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":           "NOT FOUND",
				"additional_info": err.Error(),
			})
		}

		photoUserId := photo.UserID

		log.Default().Println("id user : ", userId)

		if photoUserId != userId {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "UNAUTHORIZED, when authorizing",
			})
		}

		ctx.Next()

	}
}

func CommentAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := database.ConnectDB()
		commentRepo := repositories.NewCommentRepo(db)

		commentIdParam, err := strconv.Atoi(ctx.Param("commentID"))

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":           "BAD REQUEST, when authorizing",
				"additional_info": err.Error(),
			})
		}

		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userId := int(userData["userID"].(float64))
		log.Default().Println("userID:", userId)

		comment, err := commentRepo.GetCommentById(commentIdParam)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":           "NOT FOUND",
				"additional_info": err.Error(),
			})
		}

		commentUserId := comment.UserID

		log.Default().Println("id user : ", userId)

		if commentUserId != userId {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "UNAUTHORIZED, when authorizing",
			})
		}

		ctx.Next()

	}
}

func SocialMediaAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := database.ConnectDB()
		socialMediaRepo := repositories.NewMediaSocialRepo(db)

		socialMediaIdParam, err := strconv.Atoi(ctx.Param("socialMediaId"))

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":           "BAD REQUEST, when authorizing",
				"additional_info": err.Error(),
			})
		}

		socialMedia, err := socialMediaRepo.GetSocialMediaById(socialMediaIdParam)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":           "NOT FOUND",
				"additional_info": err.Error(),
			})
		}

		socialMediaId := socialMedia.UserID

		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userId := int(userData["userID"].(float64))
		log.Default().Println("userID:", userId)

		if socialMediaId != userId {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "UNAUTHORIZED, when authorizing",
			})
		}

		ctx.Next()

	}
}
