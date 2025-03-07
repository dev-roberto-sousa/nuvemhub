package handlers

import (
	"github.com/dev-roberto-sousa/nuvemhub/internal/api/config"
	"github.com/dev-roberto-sousa/nuvemhub/internal/api/repositories"
	"github.com/dev-roberto-sousa/nuvemhub/internal/api/services"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	config.ConnectDatabase()
	db := config.DB

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := NewUserHandler(userService)

	api := r.Group("/api/v1")
	{
		api.GET("/", HomeHandler)
		api.GET("/users", userHandler.GetUsers)
		api.POST("/users", userHandler.CreateUser)

	}

	return r
}
