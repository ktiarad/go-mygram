package server

import (
	"go-mygram/server/controllers"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	userController    *controllers.UserController
	photoController   *controllers.PhotoController
	commentController *controllers.CommentController
}

func NewServer(user *controllers.UserController, photo *controllers.PhotoController, comment *controllers.CommentController) *Server {
	return &Server{
		userController:    user,
		photoController:   photo,
		commentController: comment,
	}
}

func (s *Server) StartServer() {
	port := "8080"

	router := gin.Default()

	router.POST("/users/register", s.userController.Register)
	router.POST("/users/login", s.userController.Login)
	router.PUT("/users/:userID", s.userController.UpdateUser)
	router.DELETE("/users/:userID", s.userController.DeleteUser)

	router.POST("/photos", s.photoController.CreatePhoto)
	router.GET("/photos", s.photoController.GetAllPhotos)
	router.PUT("/photos/:userID", s.photoController.UpdatePhoto)
	router.DELETE("/photos/:userID", s.photoController.DeletePhoto)

	router.POST("/comments", s.commentController.CreateComment)
	router.GET("/comments", s.commentController.GetAllComments)
	router.PUT("/comments/:commentID", s.commentController.UpdateComment)
	router.DELETE("/comments/:commentID", s.commentController.DeleteComment)

	log.Println("Server running at port", port)

	router.Run(port)
}
