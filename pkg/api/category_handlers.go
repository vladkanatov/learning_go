package api

import (
	"net/http"
	"strconv"
	"todo-app/internal/buisness"
	"todo-app/pkg/models"

	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {
	categoryService := c.MustGet("categoryService").(*buisness.CategoryService)

	categories, err := categoryService.GetCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, categories)
}

func GetCategory(c *gin.Context) {
	categoryService := c.MustGet("categoryService").(*buisness.CategoryService)

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category, err := categoryService.GetCategory(id)
	if err != nil {
		if err.Error() == "category not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, category)
}

func CreateCategory(c *gin.Context) {
	categoryService := c.MustGet("categoryService").(*buisness.CategoryService)

	var category models.Categories
	var id int

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := categoryService.CreateCategory(&category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	category.ID = id

	c.JSON(http.StatusCreated, category)
}

func UpdateCategory(c *gin.Context) {
	categoryService := c.MustGet("categoryService").(*buisness.CategoryService)

	var category models.Categories
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category.ID = id

	if err := categoryService.UpdateCategory(&category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, category)
}

func DeleteCategory(c *gin.Context) {
	categoryService := c.MustGet("categoryService").(*buisness.CategoryService)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := categoryService.DeleteCategory(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "category has deleted"})
}
