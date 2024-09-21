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

func createTask(description string, status bool, priority int) (*http.Request, *httptest.ResponseRecorder) {
	todo := map[string]interface{}{
		"description": description,
		"status":      status,
		"priority":    priority,
	}
	jsonData, _ := json.Marshal(todo)
	req, _ := http.NewRequest("POST", "/todo", bytes.NewReader(jsonData))
	resp := httptest.NewRecorder()
	return req, resp
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

	req, resp := createTask("Getting task", true, 6)

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)

	var createdTodo map[string]interface{}
	json.Unmarshal(resp.Body.Bytes(), &createdTodo)
	id := int(createdTodo["id"].(float64))

	req, _ = http.NewRequest("GET", fmt.Sprintf("/todo/%d", id), nil)
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

	req, resp := createTask("New task", false, 5)

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)
}

func TestUpdateTodo(t *testing.T) {
	router := setupRouter()

	req, resp := createTask("Update task", false, 2)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code) // Create element

	var createdTodo map[string]interface{}
	json.Unmarshal(resp.Body.Bytes(), &createdTodo)
	id := int(createdTodo["id"].(float64))

	updatedTodo := map[string]interface{}{
		"description": "Updated task description",
		"status":      true,
		"priority":    6,
	}
	jsonData, _ := json.Marshal(updatedTodo)
	req, _ = http.NewRequest("PUT", fmt.Sprintf("/todo/%d", id), bytes.NewReader(jsonData))
	req.Header.Set("Content-Type", "application/json")
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code) // Update element

	var responseTodo map[string]interface{}
	json.Unmarshal(resp.Body.Bytes(), &responseTodo)

	assert.Equal(t, "Updated task description", responseTodo["description"])
	assert.Equal(t, true, responseTodo["status"])
	assert.Equal(t, float64(6), responseTodo["priority"])
}

func TestDeleteTodo(t *testing.T) {
	router := setupRouter()

	req, resp := createTask("Delete task", false, 5)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)

	var createdTodo map[string]interface{}
	json.Unmarshal(resp.Body.Bytes(), &createdTodo)
	id := int(createdTodo["id"].(float64))

	req, _ = http.NewRequest("DELETE", fmt.Sprintf("/todo/%d", id), nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNoContent, resp.Code)

	req, _ = http.NewRequest("GET", fmt.Sprintf("/todo/%d", id), nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)

}
