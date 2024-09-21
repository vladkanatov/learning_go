package api

import (
	"todo-app/internal/buisness"
	"todo-app/pkg/storage"

	"github.com/gin-gonic/gin"
)

func SetupRouter(db *storage.Database) *gin.Engine {
	router := gin.Default()

	todoService := buisness.NewTodoService(db)
	commentService := buisness.NewCommentService(db)

	router.Use(func(c *gin.Context) {
		c.Set("todoService", todoService)
		c.Set("commentService", commentService)
		c.Next()
	})

	router.GET("/todos", GetTodos)
	router.POST("/todo", CreateTodo)
	router.GET("/todo/:id", GetTodoById)
	router.DELETE("/todo/:id", DeleteTodo)
	router.PUT("/todo/:id", UpdateTodo)

	router.GET("/comments/:todo_id", GetCommentsByTodoID)
	router.GET("/comment/:id", GetCommentByID)
	router.POST("/comment", CreateComment)
	router.DELETE("/comment/:id", DeleteComment)
	router.PUT("/comment/:id", UpdateComment)

	return router
}
