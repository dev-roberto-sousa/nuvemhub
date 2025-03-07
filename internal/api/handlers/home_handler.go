package handlers

import (
	"net/http"

	"github.com/dev-roberto-sousa/nuvemhub/internal/api/services"
	"github.com/gin-gonic/gin"
)

func HomeHandler(c *gin.Context) {

	homeService := services.NewHomeService()

	message := homeService.GetWelcomeMessage()

	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}
