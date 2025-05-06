package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Summary Página inicial
// @Description Rota raiz da API v1
// @Tags Home
// @Produce json
// @Success 200 {object} map[string]string
// @Router /api/v1/ [get]
func HomePage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Bem vindo!"})
}
