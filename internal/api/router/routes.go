package router

import (
	"net/http"

	_ "github.com/dev-roberto-sousa/nuvemhub/docs"
	"github.com/dev-roberto-sousa/nuvemhub/internal/api/handler"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(userHandler *handler.UserHandler) *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/api/v1/", handler.HomePage)
	router.GET("/api/v1/users", userHandler.GetUsers)

	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/api/v1/")
	})

	return router
}
