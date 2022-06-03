package main

import (
	"go-mygram/database"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	port := ":8080"

	router := gin.Default()

	log.Println("Server running at port", port)

	database.ConnectDB()
	router.Run(port)
}
