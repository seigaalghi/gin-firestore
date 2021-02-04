package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/seigaalghi/firestore-go/controllers"
	middleware "github.com/seigaalghi/firestore-go/middlewares"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	server := gin.Default()

	server.GET("/api/v1/posts", middleware.AuthorizeJWT(), controllers.GetPosts)
	server.POST("/api/v1/posts", controllers.AddPost)
	server.PUT("/api/v1/posts/:id", controllers.EditPost)
	server.DELETE("/api/v1/posts/:id", controllers.DeletePost)

	server.Run(":5000")
}
