package test

import (
	"github.com/brandonromano/wrecker"
	"github.com/brandonromano/wrecker/test/models"
	"github.com/stretchr/testify/assert"
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

	err := wreckerClient.Get("/users").
		WithParam("id", "1").
		Into(&response).
		Execute()

	if err != nil {
		t.Error("Error performing GET /users")
	}

	assert.True(t, response.Success)
}

func TestFailGet(t *testing.T) {
	response := models.Response{
		Content: new(models.User),
	}

	err := wreckerClient.Get("/users").
		Into(&response).
		Execute()

	if err != nil {
		t.Error("Error performing GET /users")
	}

	assert.True(t, !response.Success)
}

func TestSuccessfulPost(t *testing.T) {
	response := models.Response{
		Content: new(models.User),
	}

	err := wreckerClient.Post("/users").
		WithParam("id", "1").
		WithParam("user_name", "BrandonRomano").
		WithParam("location", "Brooklyn, NY").
		Into(&response).
		Execute()

	if err != nil {
		t.Error("Error performing POST /users")
	}

	assert.True(t, response.Success)
}

func TestFailPost(t *testing.T) {
	response := models.Response{
		Content: new(models.User),
	}

	err := wreckerClient.Post("/users").
		WithParam("id", "1").
		WithParam("user_name", "BrandonRomano").
		Into(&response).
		Execute()

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
	err := wreckerClient.Put("/users").
		WithParam("id", "1").
		WithParam("user_name", username).
		Into(&response).
		Execute()

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

	err := wreckerClient.Put("/users").
		Into(&response).
		Execute()

	if err != nil {
		t.Error("Error performing PUT /users")
	}

	assert.True(t, !response.Success)
}

func TestSuccessfulDelete(t *testing.T) {
	response := models.Response{}

	err := wreckerClient.Delete("/users/1").
		Into(&response).
		Execute()

	if err != nil {
		t.Error("Error performing DELETE /users")
	}

	assert.True(t, response.Success)
}

func TestFailDelete(t *testing.T) {
	response := models.Response{}

	err := wreckerClient.Delete("/users/a").
		Into(&response).
		Execute()

	if err != nil {
		t.Error("Error performing DELETE /users")
	}

	assert.True(t, !response.Success)
}
