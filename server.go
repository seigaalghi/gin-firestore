package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/seigaalghi/firestore-go/controllers"
	"github.com/seigaalghi/firestore-go/middlewares"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	server := gin.Default()

	postV1 := server.Group("/api/v1/posts")
	postV1.Use(middlewares.AuthorizeJWT())

	postV1.GET("/", controllers.GetPosts)
	postV1.POST("/", controllers.AddPost)
	postV1.PUT("/:id", controllers.EditPost)
	postV1.DELETE("/:id", controllers.DeletePost)

	authV1 := server.Group("/api/v1/auth")

	authV1.POST("/register", controllers.Register)
	authV1.GET("/login", controllers.Login)

	server.Run(":5000")
}
