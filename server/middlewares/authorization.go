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
		email := string(userData["email"].(string))
		log.Default().Println("email from jwt : ", email)

		// user := models.User{}

		user := &models.User{
			Email: email,
		}

		userDb, err := userRepo.CheckUser(user)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":           "NOT FOUND, when authorizing",
				"additional_info": err.Error(),
			})
		}

		userId := userDb.ID

		log.Default().Println("id user : ", userId)

		if userIdParam != userId {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "UNAUTHORIZED, when authorizing",
			})
		}

		ctx.Next()

	}
}
