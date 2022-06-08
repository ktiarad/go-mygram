package server

import (
	"go-mygram/server/controllers"
	"go-mygram/server/middlewares"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	userController        *controllers.UserController
	photoController       *controllers.PhotoController
	commentController     *controllers.CommentController
	socialMediaController *controllers.SocialMediaController
}

func NewServer(user *controllers.UserController, photo *controllers.PhotoController, comment *controllers.CommentController, socialmedia *controllers.SocialMediaController) *Server {
	return &Server{
		userController:        user,
		photoController:       photo,
		commentController:     comment,
		socialMediaController: socialmedia,
	}
}

func (s *Server) StartServer() {
	port := ":8080"

	router := gin.Default()

	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", s.userController.Register)
		userRouter.POST("/login", s.userController.Login)
		userRouter.PUT("/:userID", middlewares.Authentication(), middlewares.UserAuthorization(), s.userController.UpdateUser)
		userRouter.DELETE("/:userID", middlewares.Authentication(), middlewares.UserAuthorization(), s.userController.DeleteUser)
	}

	photoRouter := router.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.POST("/", s.photoController.CreatePhoto)
		photoRouter.GET("/", s.photoController.GetAllPhotos)
		photoRouter.PUT("/:photoID", middlewares.PhotoAuthorization(), s.photoController.UpdatePhoto)
		photoRouter.DELETE("/:photoID", middlewares.PhotoAuthorization(), s.photoController.DeletePhoto)
	}

	commentRouter := router.Group("/comments")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.POST("/", s.commentController.CreateComment)
		commentRouter.GET("/", s.commentController.GetAllComments)
		commentRouter.PUT("/:commentID", middlewares.CommentAuthorization(), s.commentController.UpdateComment)
		commentRouter.DELETE("/:commentID", middlewares.CommentAuthorization(), s.commentController.DeleteComment)
	}

	socialMediaRouter := router.Group("/socialmedias")
	{
		socialMediaRouter.Use(middlewares.Authentication())
		socialMediaRouter.POST("/", s.socialMediaController.CreateSocialMedia)
		socialMediaRouter.GET("/", s.socialMediaController.GetAllSocialMedias)
		socialMediaRouter.PUT("/:socialMediaId", middlewares.SocialMediaAuthorization(), s.socialMediaController.UpdateSocialMedia)
		socialMediaRouter.DELETE("/:socialMediaId", middlewares.SocialMediaAuthorization(), s.socialMediaController.DeleteSocialMedia)
	}

	log.Println("Server running at port", port)

	router.Run(port)
}
