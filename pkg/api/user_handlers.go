package api

import (
	"todo-app/internal/buisness"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	userService := c.MustGet("userService").(*buisness.UserService)

}
