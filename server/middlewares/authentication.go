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
				// })
				// return &params.Response{
				// 	Status:         http.StatusUnauthorized,
				// 	Error:          "UNAUTHENTICATED",
				// 	AdditionalInfo: err.Error(),
				// }
			})
		}

		ctx.Set("userData", verifyToken)
		ctx.Next()
		// return func(c *gin.Context) {
		// 	verifyToken, err := helpers.VerifyToken(c)
		// 	_ = verifyToken

		// 	if err != nil {

		// 	}
		// }
	}

}
