package api

import (
	"net/http"
	"strconv"
	"todo-app/internal/buisness"
	"todo-app/pkg/models"

	"github.com/gin-gonic/gin"
)

func GetCommentByID(c *gin.Context) {
	commentService := c.MustGet("commentService").(*buisness.CommentService)

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment, err := commentService.GetCommentByID(id)
	if err != nil {
		if err.Error() == "comment not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, comment)
}

func GetCommentsByTodoID(c *gin.Context) {
	commentService := c.MustGet("commentService").(*buisness.CommentService)

	todoID, err := strconv.Atoi(c.Param("todo_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comments, err := commentService.GetCommentsByTodoID(todoID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, comments)
}

func CreateComment(c *gin.Context) {
	commentService := c.MustGet("commentService").(*buisness.CommentService)

	var comment models.Comment

	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, created_at, err := commentService.CreateComment(&comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	comment.ID = id
	comment.CreatedAt = created_at

	c.JSON(http.StatusCreated, comment)
}

func DeleteComment(c *gin.Context) {
	commentService := c.MustGet("commentService").(*buisness.CommentService)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := commentService.DeleteCommentByID(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "comment was deleted"})
}

func UpdateComment(c *gin.Context) {
	commentService := c.MustGet("commentService").(*buisness.CommentService)

	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	comment.ID = id
	if err := commentService.UpdateCommentByID(&comment); err != nil {
		if err.Error() == "comment not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, comment)
}
