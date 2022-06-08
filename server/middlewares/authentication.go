package middlewares

import (
	"go-mygram/helpers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		verifyToken, err := helpers.VerifyToken(ctx)
		log.Default().Println("Verify token :", verifyToken)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":           "UNAUTHENTICATED",
				"additional_info": err.Error(),
			})
		}

		ctx.Set("userData", verifyToken)
		ctx.Next()
	}

}
