package test

import (
	"net/http"
	"os"
	"testing"

	"github.com/BrandonRomano/wrecker"
	"github.com/BrandonRomano/wrecker/test/models"
	"github.com/stretchr/testify/assert"
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
		URLParam("id", "1").
		Into(&response).
		Execute()

	assert.Nil(t, err)
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

	assert.NotNil(t, err)
	assert.Equal(t, http.StatusBadRequest, httpResponse.StatusCode)
}

func TestSuccessfulPost(t *testing.T) {
	response := models.Response{
		Content: new(models.User),
	}

	httpResponse, err := wreckerClient.Post("/users").
		FormParam("id", "1").
		FormParam("user_name", "BrandonRomano").
		FormParam("location", "Brooklyn, NY").
		Into(&response).
		Execute()

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, httpResponse.StatusCode)
	assert.Equal(t, response.Content.(*models.User).UserName, "BrandonRomano")
}

func TestFailPost(t *testing.T) {
	response := models.Response{
		Content: new(models.User),
	}

	httpResponse, err := wreckerClient.Post("/users").
		FormParam("id", "1").
		FormParam("user_name", "BrandonRomano").
		Into(&response).
		Execute()

	if err == nil {
		t.Error("Error expected for POST /users")
	}

	assert.Equal(t, http.StatusBadRequest, httpResponse.StatusCode)
}

func TestSuccessfulPut(t *testing.T) {
	response := models.Response{
		Content: new(models.User),
	}

	username := "BrandonRomano100"
	httpResponse, err := wreckerClient.Put("/users").
		FormParam("id", "1").
		FormParam("user_name", username).
		Into(&response).
		Execute()

	assert.Nil(t, err)
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

	assert.NotNil(t, err)
	assert.Equal(t, http.StatusBadRequest, httpResponse.StatusCode)
}

func TestSuccessfulDelete(t *testing.T) {
	response := models.Response{}

	httpResponse, err := wreckerClient.Delete("/users/1").
		Header("delete-test-header", "delete-test-header-value").
		Into(&response).
		Execute()

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, httpResponse.StatusCode)
}

func TestDeleteFailFromURL(t *testing.T) {
	response := models.Response{}

	httpResponse, err := wreckerClient.Delete("/users/a").
		Header("delete-test-header", "delete-test-header-value").
		Into(&response).
		Execute()

	assert.NotNil(t, err)
	assert.Equal(t, http.StatusBadRequest, httpResponse.StatusCode)
}

func TestDeleteFailFromHeader(t *testing.T) {
	response := models.Response{}

	httpResponse, err := wreckerClient.Delete("/users/1").
		Into(&response).
		Execute()

	assert.NotNil(t, err)
	assert.Equal(t, http.StatusBadRequest, httpResponse.StatusCode)
}
