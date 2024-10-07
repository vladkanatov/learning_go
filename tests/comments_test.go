package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createComment(todoID int, content string, author string) (*http.Request, *httptest.ResponseRecorder) {
	comment := map[string]interface{}{
		"todo_id": todoID,
		"content": content,
		"author":  author,
	}
	jsonData, _ := json.Marshal(&comment)
	req, _ := http.NewRequest("POST", "/comment", bytes.NewReader(jsonData))
	resp := httptest.NewRecorder()
	return req, resp
}

func TestCreateComment(t *testing.T) {
	router := setupRouter()

	req, resp := createComment(1, "Привет, я Влад", "Vladislav Kanatov")
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)
}

func TestGetCommentByID(t *testing.T) {
	router := setupRouter()

	req, resp := createComment(6, "Привет, я Владик-шоколадик", "Vladislav Kanatov")
	router.ServeHTTP(resp, req)

	var createdComment map[string]interface{}
	json.Unmarshal(resp.Body.Bytes(), &createdComment)
	id := int(createdComment["id"].(float64))

	req, _ = http.NewRequest("GET", fmt.Sprintf("/comment/%d", id), nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	var responseComment map[string]interface{}
	json.Unmarshal(resp.Body.Bytes(), &responseComment)

	assert.Equal(t, float64(6), responseComment["todo_id"])
	assert.Equal(t, "Привет, я Владик-шоколадик", responseComment["content"])
	assert.Equal(t, "Vladislav Kanatov", responseComment["author"])
}

func TestGetCommentsByTodoID(t *testing.T) {
	router := setupRouter()

	req, resp := createComment(2, "Some Tests", "Vladislav Kanatov")
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)

	var createdComment map[string]interface{}
	json.Unmarshal(resp.Body.Bytes(), &createdComment)

	todoID := int(createdComment["todo_id"].(float64))

	req, _ = http.NewRequest("GET", fmt.Sprintf("/comments/%d", todoID), nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestDeleteComment(t *testing.T) {
	router := setupRouter()

	req, resp := createComment(3, "Test for delete", "Vladislav Kanatov")
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)

	var createdComment map[string]interface{}
	json.Unmarshal(resp.Body.Bytes(), &createdComment)
	id := int(createdComment["id"].(float64))

	req, _ = http.NewRequest("DELETE", fmt.Sprintf("/comment/%d", id), nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	req, _ = http.NewRequest("GET", fmt.Sprintf("/comment/%d", id), nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
}

func TestUpdateComment(t *testing.T) {
	router := setupRouter()

	req, resp := createComment(10, "Update comment", "Vladislav Kanatov")
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)

	var createdComment map[string]interface{}
	json.Unmarshal(resp.Body.Bytes(), &createdComment)
	id := int(createdComment["id"].(float64))

	updatedComment := map[string]interface{}{
		"todo_id": 10,
		"content": "Updated comment lol kek cheburek",
		"author":  "Vladislav Kanatov",
	}

	jsonData, _ := json.Marshal(updatedComment)

	req, _ = http.NewRequest("PUT", fmt.Sprintf("/comment/%d", id), bytes.NewReader(jsonData))
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	var responseComment map[string]interface{}
	json.Unmarshal(resp.Body.Bytes(), &responseComment)

	assert.Equal(t, float64(10), responseComment["todo_id"])
	assert.Equal(t, "Updated comment lol kek cheburek", responseComment["content"])
	assert.Equal(t, "Vladislav Kanatov", responseComment["author"])
}
