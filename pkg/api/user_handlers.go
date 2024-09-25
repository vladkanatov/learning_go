package api

import (
	"net/http"
	"todo-app/internal/buisness"
	"todo-app/pkg/models"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	userService := c.MustGet("userService").(*buisness.UserService)

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := userService.LoginUser(&user); err != nil {
		if err.Error() == "invalid password" {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successful login!"})
}
