package test

import (
	"github.com/brandonromano/wrecker"
	"github.com/brandonromano/wrecker/test/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

var wreckerClient *wrecker.Wrecker

func init() {
	go startServer()
	wreckerClient = wrecker.New("http://localhost:" + os.Getenv("PORT"))
}

func TestSuccessfulGet(t *testing.T) {
	response := models.Response{
		Content: new(models.User),
	}

	httpResponse, err := wreckerClient.Get("/users").
		WithParam("id", "1").
		Into(&response).
		Execute()

	if err != nil {
		t.Error("Error performing GET /users")
	}

	assert.Equal(t, http.StatusOK, httpResponse.StatusCode)
	assert.Equal(t, response.Content.(*models.User).UserName, "BrandonRomano")
}

func TestFailGet(t *testing.T) {
	response := models.Response{
		Content: new(models.User),
	}

	httpResponse, err := wreckerClient.Get("/users").
		Into(&response).
		Execute()

	if err != nil {
		t.Error("Error performing GET /users")
	}

	assert.Equal(t, http.StatusBadRequest, httpResponse.StatusCode)
}

func TestSuccessfulPost(t *testing.T) {
	response := models.Response{
		Content: new(models.User),
	}

	httpResponse, err := wreckerClient.Post("/users").
		WithParam("id", "1").
		WithParam("user_name", "BrandonRomano").
		WithParam("location", "Brooklyn, NY").
		Into(&response).
		Execute()

	if err != nil {
		t.Error("Error performing POST /users")
	}

	assert.Equal(t, http.StatusOK, httpResponse.StatusCode)
	assert.Equal(t, response.Content.(*models.User).UserName, "BrandonRomano")
}

func TestFailPost(t *testing.T) {
	response := models.Response{
		Content: new(models.User),
	}

	httpResponse, err := wreckerClient.Post("/users").
		WithParam("id", "1").
		WithParam("user_name", "BrandonRomano").
		Into(&response).
		Execute()

	if err != nil {
		t.Error("Error performing POST /users")
	}

	assert.Equal(t, http.StatusBadRequest, httpResponse.StatusCode)
}

func TestSuccessfulPut(t *testing.T) {
	response := models.Response{
		Content: new(models.User),
	}

	username := "BrandonRomano100"
	httpResponse, err := wreckerClient.Put("/users").
		WithParam("id", "1").
		WithParam("user_name", username).
		Into(&response).
		Execute()

	if err != nil {
		t.Error("Error performing PUT /users")
	}

	assert.Equal(t, http.StatusOK, httpResponse.StatusCode)
	assert.Equal(t, response.Content.(*models.User).UserName, username)
}

func TestFailPut(t *testing.T) {
	response := models.Response{
		Content: new(models.User),
	}

	httpResponse, err := wreckerClient.Put("/users").
		Into(&response).
		Execute()

	if err != nil {
		t.Error("Error performing PUT /users")
	}

	assert.Equal(t, http.StatusBadRequest, httpResponse.StatusCode)
}

func TestSuccessfulDelete(t *testing.T) {
	response := models.Response{}

	httpResponse, err := wreckerClient.Delete("/users/1").
		WithHeader("delete-test-header", "delete-test-header-value").
		Into(&response).
		Execute()

	if err != nil {
		t.Error("Error performing DELETE /users")
	}

	assert.Equal(t, http.StatusOK, httpResponse.StatusCode)
}

func TestDeleteFailFromURL(t *testing.T) {
	response := models.Response{}

	httpResponse, err := wreckerClient.Delete("/users/a").
		WithHeader("delete-test-header", "delete-test-header-value").
		Into(&response).
		Execute()

	if err != nil {
		t.Error("Error performing DELETE /users")
	}

	assert.Equal(t, http.StatusBadRequest, httpResponse.StatusCode)
}

func TestDeleteFailFromHeader(t *testing.T) {
	response := models.Response{}

	httpResponse, err := wreckerClient.Delete("/users/1").
		Into(&response).
		Execute()

	if err != nil {
		t.Error("Error performing DELETE /users")
	}

	assert.Equal(t, http.StatusBadRequest, httpResponse.StatusCode)
}
