package main

import (
	"redis_sample_project/config"
	"redis_sample_project/controllers"
	"redis_sample_project/models"

	"github.com/gin-gonic/gin"
)

func main() {

	// Initialize PostgresSQL and Redis connections
	config.InitDB()
	config.InitRedis()

	// Auto-migrate User model
	config.DB.AutoMigrate(&models.User{})

	r := gin.Default()

	// Routes
	r.POST("/users", controllers.CreateUser)
	r.GET("/users/:id", controllers.GetUserByID)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)
	r.GET("/users", controllers.GetAllUsers)

	// Start the server
	r.Run(":8080")
}
