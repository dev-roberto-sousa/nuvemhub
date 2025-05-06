package main

import (
	"github.com/dev-roberto-sousa/nuvemhub/internal/api/config"
	"github.com/dev-roberto-sousa/nuvemhub/internal/api/handler"
	"github.com/dev-roberto-sousa/nuvemhub/internal/api/repository"
	"github.com/dev-roberto-sousa/nuvemhub/internal/api/router"
	"github.com/dev-roberto-sousa/nuvemhub/internal/api/service"
)

func main() {
	config.ConnectDatabase()

	userRepo := repository.NewUserRepository(config.DB)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	r := router.SetupRouter(userHandler)

	r.Run()
}
