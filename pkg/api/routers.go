package api

import (
	"todo-app/internal/buisness"
	"todo-app/pkg/storage"

	"github.com/gin-gonic/gin"
)

func SetupRouter(db *storage.Database) *gin.Engine {
	router := gin.Default()

	todoService := buisness.NewTodoService(db)

	router.Use(func(c *gin.Context) {
		c.Set("todoService", todoService)
		c.Next()
	})

	router.GET("/todos", GetTodos)
	router.POST("/todos", CreateTodo)
	router.GET("/todos/:id", GetTodoById)
	router.DELETE("/todos/:id", DeleteTodo)
	router.PUT("/todos/:id", UpdateTodo)

	return router
}
