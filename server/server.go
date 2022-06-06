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
		userRouter.DELETE("/:userID", s.userController.DeleteUser)
	}
	// router.POST("/users/register", s.userController.Register)
	// router.POST("/users/login", s.userController.Login)
	// router.PUT("/users/:userID", s.userController.UpdateUser)
	// router.DELETE("/users/:userID", s.userController.DeleteUser)

	photoRouter := router.Group("/photos")
	{
		photoRouter.POST("/", s.photoController.CreatePhoto)
		photoRouter.GET("/", s.photoController.GetAllPhotos)
		photoRouter.PUT("/:userID", s.photoController.UpdatePhoto)
		photoRouter.DELETE("/:userID", s.photoController.DeletePhoto)
	}

	commentRouter := router.Group("/comments")
	{
		commentRouter.POST("/", s.commentController.CreateComment)
		commentRouter.GET("/", s.commentController.GetAllComments)
		commentRouter.PUT("/:commentID", s.commentController.UpdateComment)
		commentRouter.DELETE("/:commentID", s.commentController.DeleteComment)
	}

	// router.POST("/comments", s.commentController.CreateComment)
	// router.GET("/comments", s.commentController.GetAllComments)
	// router.PUT("/comments/:commentID", s.commentController.UpdateComment)
	// router.DELETE("/comments/:commentID", s.commentController.DeleteComment)

	socialMediaRouter := router.Group("/socialmedias")
	{
		socialMediaRouter.POST("/", s.socialMediaController.CreateSocialMedia)
		socialMediaRouter.GET("/", s.socialMediaController.GetAllSocialMedias)
		socialMediaRouter.PUT("/:socialmediaID", s.socialMediaController.UpdateSocialMedia)
		socialMediaRouter.DELETE("/:socialmediaID", s.socialMediaController.DeleteSocialMedia)
	}
	// router.POST("/socialmedias", s.socialMediaController.CreateSocialMedia)
	// router.GET("/socialmedias", s.socialMediaController.GetAllSocialMedias)
	// router.PUT("/socialmedias/:socialmediaID", s.socialMediaController.UpdateSocialMedia)
	// router.DELETE("/socialmedias/:socialmediaID", s.socialMediaController.DeleteSocialMedia)

	log.Println("Server running at port", port)

	router.Run(port)
}
