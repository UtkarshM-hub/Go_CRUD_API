package main

import (
	"github.com/Go_CRUD_API/Routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main(){
	// loading all the environment variables so that they are accessible
	godotenv.Load(".env")

	router:=gin.Default()

	router.POST("/",routes.CreatePost)
	router.GET("/post/:postId",routes.GetOne)

	// Connecting to the database
	router.Run("localhost:8080")
}