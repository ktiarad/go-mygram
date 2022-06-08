package main

import (
	"go-mygram/database"
	"go-mygram/repositories"
	"go-mygram/server"
	"go-mygram/server/controllers"
	"go-mygram/services"
)

func main() {

	db := database.ConnectDB()
	userRepo := repositories.NewUserRepo(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	photoRepo := repositories.NewPhotoRepo(db)
	photoService := services.NewPhotoService(photoRepo)
	photoController := controllers.NewPhotoController(photoService)

	commentRepo := repositories.NewCommentRepo(db)
	commentService := services.NewCommentService(commentRepo)
	commentController := controllers.NewCommentController(commentService)

	socialMediaRepo := repositories.NewMediaSocialRepo(db)
	socialMediaService := services.NewSocialMediaService(socialMediaRepo)
	socialMediaController := controllers.NewSocialMediaController(socialMediaService)

	s := server.NewServer(userController, photoController, commentController, socialMediaController)
	s.StartServer()
}
