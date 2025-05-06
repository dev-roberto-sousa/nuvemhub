package handler

import (
	"net/http"

	"github.com/dev-roberto-sousa/nuvemhub/internal/api/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	svc *service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

// @Summary Listar usuários
// @Description Retorna todos os usuários cadastrados
// @Tags Users
// @Produce json
// @Success 200 {object} map[string]string
// @Router /api/v1/users [get]
func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.svc.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar usuários."})

		return
	}

	c.JSON(http.StatusOK, users)
}
