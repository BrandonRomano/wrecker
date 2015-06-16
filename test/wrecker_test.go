package test

import (
	"github.com/brandonromano/wrecker"
	"github.com/brandonromano/wrecker/test/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"os"
	"testing"
	"time"
)

var wreckerClient wrecker.Wrecker

func init() {
	go startServer()

	wreckerClient = wrecker.Wrecker{
		BaseURL: "http://localhost:" + os.Getenv("PORT"),
		HttpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func TestSuccessfulGet(t *testing.T) {
	params := url.Values{}
	params.Add("id", "1")

	response := models.Response{
		Content: new(models.User),
	}

	err := wreckerClient.Get("/users", params, &response)
	if err != nil {
		t.Error("Error performing GET /users")
	}

	assert.True(t, response.Success)
}

func TestFailGet(t *testing.T) {
	response := models.Response{
		Content: new(models.User),
	}

	err := wreckerClient.Get("/users", nil, &response)
	if err != nil {
		t.Error("Error performing GET /users")
	}

	assert.True(t, !response.Success)
}

func TestSuccessfulPost(t *testing.T) {
	response := models.Response{
		Content: new(models.User),
	}

	params := url.Values{}
	params.Add("id", "1")
	params.Add("user_name", "BrandonRomano")
	params.Add("location", "Brooklyn, NY")

	err := wreckerClient.Post("/users", params, &response)
	if err != nil {
		t.Error("Error performing POST /users")
	}

	assert.True(t, response.Success)
}

func TestFailPost(t *testing.T) {
	response := models.Response{
		Content: new(models.User),
	}

	params := url.Values{}
	params.Add("id", "1")
	params.Add("user_name", "BrandonRomano")

	err := wreckerClient.Post("/users", params, &response)
	if err != nil {
		t.Error("Error performing POST /users")
	}

	assert.True(t, !response.Success)
}

func TestSuccessfulPut(t *testing.T) {
	response := models.Response{
		Content: new(models.User),
	}

	username := "BrandonRomano100"
	params := url.Values{}
	params.Add("id", "1")
	params.Add("user_name", username)

	err := wreckerClient.Put("/users", params, &response)
	if err != nil {
		t.Error("Error performing PUT /users")
	}

	assert.True(t, response.Success)

	user := response.Content.(*models.User)
	assert.Equal(t, user.UserName, username)
}

func TestFailPut(t *testing.T) {
	response := models.Response{
		Content: new(models.User),
	}

	err := wreckerClient.Put("/users", nil, &response)
	if err != nil {
		t.Error("Error performing PUT /users")
	}

	assert.True(t, !response.Success)
}

func TestSuccessfulDelete(t *testing.T) {
	response := models.Response{}

	err := wreckerClient.Delete("/users/1", &response)
	if err != nil {
		t.Error("Error performing DELETE /users")
	}

	assert.True(t, response.Success)
}

func TestFailDelete(t *testing.T) {
	response := models.Response{}

	err := wreckerClient.Delete("/users/a", &response)
	if err != nil {
		t.Error("Error performing DELETE /users")
	}

	assert.True(t, !response.Success)
}
