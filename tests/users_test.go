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

func createUser(username, email, password string) (*http.Request, *httptest.ResponseRecorder) {
	user := map[string]interface{}{
		"username": username,
		"email":    email,
		"password": password,
	}
	jsonData, _ := json.Marshal(&user)
	req, _ := http.NewRequest("POST", "/user/", bytes.NewReader(jsonData))
	resp := httptest.NewRecorder()

	return req, resp

}

func TestCreateUser(t *testing.T) {
	router := setupRouter()

	req, resp := createUser("vladkanatov", "vladkanatov@email.com", "132465-Cs")

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)

}

func TestGetUser(t *testing.T) {
	router := setupRouter()

	req, resp := createUser("kanatovlad", "vlad.kanaqta@ya.ru", "1fdsavx")

	router.ServeHTTP(resp, req)

	var createdUser map[string]interface{}
	json.Unmarshal(resp.Body.Bytes(), &createdUser)

	id := int(createdUser["id"].(float64))

	req, _ = http.NewRequest("GET", fmt.Sprintf("/user/%d", id), nil)
	resp = httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestGetUsers(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/users", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestUpdate(t *testing.T) {
	router := setupRouter()

	req, resp := createUser("timofeyka", "timofeyka@ya.ru", "favfals")

	router.ServeHTTP(resp, req)

	var createdUser map[string]interface{}
	json.Unmarshal(resp.Body.Bytes(), &createdUser)

	id := int(createdUser["id"].(float64))

	updateUser := map[string]interface{}{
		"username": "podofeyka",
		"email":    "podofeyka@email.com",
		"password": "faxgaasf",
	}
	jsonData, _ := json.Marshal(&updateUser)

	req, _ = http.NewRequest("PUT", fmt.Sprintf("/user/%d", id), bytes.NewReader(jsonData))
	resp = httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	var responseUser map[string]interface{}
	json.Unmarshal(resp.Body.Bytes(), &responseUser)

	assert.Equal(t, updateUser["username"], responseUser["username"])
	assert.Equal(t, updateUser["email"], responseUser["email"])
	assert.Equal(t, updateUser["password"], responseUser["password"])
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestDelete(t *testing.T) {
	router := setupRouter()

	req, resp := createUser("sakas", "sakas@ya.ru", "sdafzxva")

	router.ServeHTTP(resp, req)

	var createdUser map[string]interface{}
	json.Unmarshal(resp.Body.Bytes(), &createdUser)

	id := int(createdUser["id"].(float64))

	req, _ = http.NewRequest("DELETE", fmt.Sprintf("/user/%d", id), nil)
	resp = httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	req, _ = http.NewRequest("GET", fmt.Sprintf("/user/%d", id), nil)
	resp = httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
}
