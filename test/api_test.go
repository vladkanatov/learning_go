package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo-app/pkg/api"
	"todo-app/pkg/storage"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	db := storage.SetupDatabase()
	return api.SetupRouter(db)
}

func TestGetTodos(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/todos", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestGetTodoById(t *testing.T) {
	router := setupRouter()

	todo := map[string]interface{}{
		"description": "Getting task",
		"status":      true,
		"priority":    6,
	}
	jsonData, _ := json.Marshal(todo)
	req, _ := http.NewRequest("POST", "/todos", bytes.NewReader(jsonData))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)

	var createdTodo map[string]interface{}
	json.Unmarshal(resp.Body.Bytes(), &createdTodo)
	id := int(createdTodo["id"].(float64))

	req, _ = http.NewRequest("GET", fmt.Sprintf("/todos/%d", id), nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	var responseTodo map[string]interface{}
	json.Unmarshal(resp.Body.Bytes(), &responseTodo)

	assert.Equal(t, "Getting task", responseTodo["description"])
	assert.Equal(t, true, responseTodo["status"])
	assert.Equal(t, float64(6), responseTodo["priority"])

}

func TestCreateTodo(t *testing.T) {
	router := setupRouter()

	todo := map[string]interface{}{
		"description": "new task",
		"status":      false,
		"priority":    5,
	}

	jsonData, _ := json.Marshal(todo)

	req, _ := http.NewRequest("POST", "/todos", bytes.NewReader(jsonData))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)
}

func TestUpdateTodo(t *testing.T) {
	router := setupRouter()

	todo := map[string]interface{}{
		"description": "Update task",
		"status":      false,
		"priority":    2,
	}
	jsonData, _ := json.Marshal(todo)
	req, _ := http.NewRequest("POST", "/todos", bytes.NewReader(jsonData))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)

	var createdTodo map[string]interface{}
	json.Unmarshal(resp.Body.Bytes(), &createdTodo)
	id := int(createdTodo["id"].(float64))

	updatedTodo := map[string]interface{}{
		"description": "Updated task description",
		"status":      true,
		"priority":    6,
	}
	jsonData, _ = json.Marshal(updatedTodo)
	req, _ = http.NewRequest("PUT", fmt.Sprintf("/todos/%d", id), bytes.NewReader(jsonData))
	req.Header.Set("Content-Type", "application/json")
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var responseTodo map[string]interface{}
	json.Unmarshal(resp.Body.Bytes(), &responseTodo)

	assert.Equal(t, "Updated task description", responseTodo["description"])
	assert.Equal(t, true, responseTodo["status"])
	assert.Equal(t, float64(6), responseTodo["priority"])
}

func TestDeleteTodo(t *testing.T) {
	router := setupRouter()

	todo := map[string]interface{}{
		"description": "Delete task",
		"status":      false,
		"priority":    5,
	}

	jsonData, _ := json.Marshal(todo)
	req, _ := http.NewRequest("POST", "/todos", bytes.NewReader(jsonData))

	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)

	var createdTodo map[string]interface{}
	json.Unmarshal(resp.Body.Bytes(), &createdTodo)
	id := int(createdTodo["id"].(float64))

	req, _ = http.NewRequest("DELETE", fmt.Sprintf("/todos/%d", id), nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNoContent, resp.Code)

	req, _ = http.NewRequest("GET", fmt.Sprintf("/todos/%d", id), nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)

}
