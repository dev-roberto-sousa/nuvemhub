package router

import (
	"github.com/dev-roberto-sousa/nuvemhub/internal/api/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter(userHandler *handler.UserHandler) *gin.Engine {
	router := gin.Default()

	router.GET("api/v1/", handler.HomePage)
	router.GET("api/v1/users", userHandler.GetUsers)

	return router
}
