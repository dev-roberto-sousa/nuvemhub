package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	users := []string{"user1", "user2", "user3"}

	c.JSON(http.StatusOK, users)
}
