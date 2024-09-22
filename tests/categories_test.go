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

func createCategory(name string) (*http.Request, *httptest.ResponseRecorder) {

	category := map[string]interface{}{
		"name": name,
	}

	jsonData, _ := json.Marshal(&category)

	req, _ := http.NewRequest("POST", "/category", bytes.NewReader(jsonData))
	resp := httptest.NewRecorder()

	return req, resp
}

func TestCreateCategory(t *testing.T) {
	router := setupRouter()

	req, resp := createCategory("DevOps")

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)
}

func TestGetCategory(t *testing.T) {
	router := setupRouter()

	req, resp := createCategory("Qwer")

	router.ServeHTTP(resp, req)

	var createdCategory map[string]interface{}
	json.Unmarshal(resp.Body.Bytes(), &createdCategory)
	id := int(createdCategory["id"].(float64))

	req, _ = http.NewRequest("GET", fmt.Sprintf("/category/%d", id), nil)
	resp = httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	var responseCategory map[string]interface{}
	json.Unmarshal(resp.Body.Bytes(), &responseCategory)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, responseCategory["id"], createdCategory["id"])
	assert.Equal(t, responseCategory["name"], createdCategory["name"])
}

func TestGetCategories(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/categories", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestUpdateCategory(t *testing.T) {
	router := setupRouter()

	req, resp := createCategory("Frontend")
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(resp, req)

	var createdCategory map[string]interface{}
	json.Unmarshal(resp.Body.Bytes(), &createdCategory)
	id := int(createdCategory["id"].(float64))

	updatedCategory := map[string]interface{}{
		"name": "Backend",
	}

	jsonData, _ := json.Marshal(&updatedCategory)
	req, _ = http.NewRequest("PUT", fmt.Sprintf("/category/%d", id), bytes.NewReader(jsonData))
	resp = httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	var responseCategory map[string]interface{}
	json.Unmarshal(resp.Body.Bytes(), &responseCategory)

	assert.Equal(t, updatedCategory["name"], responseCategory["name"])
	assert.Equal(t, float64(id), responseCategory["id"])
	assert.Equal(t, http.StatusOK, resp.Code)

}

func TestDeleteCategory(t *testing.T) {
	router := setupRouter()

	req, resp := createCategory("UX/UI")
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(resp, req)

	var createdCategory map[string]interface{}
	json.Unmarshal(resp.Body.Bytes(), &createdCategory)

	id := int(createdCategory["id"].(float64))

	req, _ = http.NewRequest("DELETE", fmt.Sprintf("/category/%d", id), nil)
	resp = httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	req, _ = http.NewRequest("GET", fmt.Sprintf("/category/%d", id), nil)
	resp = httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
}
