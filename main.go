package main

import (
	configs "user-crud/config"
	"user-crud/controllers"
	middleware "user-crud/middlewares"
	"user-crud/seed"

	"github.com/gin-gonic/gin"
)

func main() {
	db := configs.InitDB()
	seed.SeedMasterAdmin(db)

	router := gin.Default()

	api := router.Group("/api")
	{
		api.POST("/register", controllers.RegisterUser)
		api.POST("/login", controllers.LoginUser)
	}

	authApi := router.Group("/api/user")
	authApi.Use(middleware.AuthMiddleware())
	{
		authApi.GET("/:id", controllers.GetUserByID)
		authApi.PUT("/:id", controllers.UpdateUser)
		authApi.DELETE("/:id", controllers.DeleteUser)
		authApi.GET("/list", controllers.GetUserList)
	}

	router.Run(":8080")
}
