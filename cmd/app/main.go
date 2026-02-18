package main

import (
	"github.com/gin-gonic/gin"
	// swaggerFiles "github.com/swaggo/files"
	// ginSwagger "github.com/swaggo/gin-swagger"

	"base/api"
	_ "base/docs"
	"base/repository"
	"base/config"
	"fmt"
)

func main() {

	config.LoadConfig()

	repository.ConnectDatabase()
	
	r := gin.Default()

	authGroup := r.Group("/auth")
	{
		authGroup.POST("/register", api.Register)
		authGroup.POST("/login", api.Login)
	}

	protected := r.Group("/api")
	protected.Use(api.AuthMiddleware())
	{
		protected.GET("/me", func(c *gin.Context) {
			userId, _ := c.Get("user_id")
			c.JSON(200, gin.H{"your_id": userId})
		})
		protected.POST("/events", api.CreateEvent)
		protected.GET("/events", api.GetEvents)
		protected.GET("/events/:id", api.GetEventByID)
		protected.PUT("/events/:id", api.UpdateEvent)
		protected.DELETE("/events/:id", api.DeleteEvent)
		protected.POST("/events/:id/apply", api.ApplyToEvent)
		protected.GET("my/events", api.GetMyRegistrations)
	}
	
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	fmt.Println("Server running on port 8080")
	r.Run(":8080")
}