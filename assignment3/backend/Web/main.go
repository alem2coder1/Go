package main

import (
	"backend/Web/apihelper"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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
	router.GET("/api/users", apihelper.GetAllUsersHandler)
	router.POST("/api/users", apihelper.AddUserHandler)
	router.GET("/api/users/:id", apihelper.GetUserHandler)
	router.PUT("/api/users/:id", apihelper.UpdateUserHandler)
	router.DELETE("/api/users/:id", apihelper.DeleteUserHandler)
	router.POST("/api/users/login", apihelper.LoginHandler)

	router.Run(":8083")
}
