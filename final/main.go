package main

import (
	"final/Common"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	router := gin.Default()

	corsConfig := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}
	router.Use(cors.New(corsConfig))

	api := router.Group("/api/v1")
	{
		users := api.Group("/users")
		{
			users.GET("", apihelper.GetAllUsersHandler)
			users.POST("", apihelper.AddUserHandler)
			users.GET("/:id", apihelper.GetUserHandler)
			users.PUT("/:id", apihelper.UpdateUserHandler)
			users.DELETE("/:id", apihelper.DeleteUserHandler)
			users.POST("/login", apihelper.LoginHandler)
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8083"
	}

	if err := router.Run(":" + port); err != nil {
		panic("Failed to start the server: " + err.Error())
	}
}
