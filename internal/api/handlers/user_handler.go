package handlers

import (
	"net/http"

	"github.com/dev-roberto-sousa/nuvemhub/internal/api/models"
	"github.com/dev-roberto-sousa/nuvemhub/internal/api/services"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService services.IUserService
}

func NewUserHandler(userService services.IUserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to search for users.",
			"details": err.Error(),
		})
		return
	}

	usersPublic := make([]models.UserPublic, 0, len(users))

	for _, user := range users {
		usersPublic = append(usersPublic, models.UserPublic{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		})
	}

	c.JSON(http.StatusOK, gin.H{"users": usersPublic})
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request body.",
			"details": err.Error(),
		})
		return
	}

	createdUser, err := h.userService.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create user.",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": createdUser})
}
