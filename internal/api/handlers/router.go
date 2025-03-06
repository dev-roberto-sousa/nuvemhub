package handlers

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", HomeHandler)

	return r
}
