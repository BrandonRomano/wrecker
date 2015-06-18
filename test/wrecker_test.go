package test

import (
	"github.com/benpate/wrecker"
	"github.com/benpate/wrecker/test/models"
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
		WithHeader("delete-test-header", "delete-test-header-value").
		Into(&response).
		Execute()

	if err != nil {
		t.Error("Error performing DELETE /users")
	}

	assert.True(t, response.Success)
}

func TestDeleteFailFromURL(t *testing.T) {
	response := models.Response{}

	err := wreckerClient.Delete("/users/a").
		WithHeader("delete-test-header", "delete-test-header-value").
		Into(&response).
		Execute()

	if err != nil {
		t.Error("Error performing DELETE /users")
	}

	assert.True(t, !response.Success)
}

func TestDeleteFailFromHeader(t *testing.T) {
	response := models.Response{}

	err := wreckerClient.Delete("/users/1").
		Into(&response).
		Execute()

	if err != nil {
		t.Error("Error performing DELETE /users")
	}

	assert.True(t, !response.Success)
}

func TestRestPost(t *testing.T) {

	response := models.Response{
		Content: new(models.User),
	}

	userIn := models.User{
		Id:       98,
		UserName: "Steve Rogers",
		Location: "New York, NY",
	}

	err := wreckerClient.Post("/users").
		WithBody(userIn).
		Into(&response).
		Execute()

	assert.True(t, err == nil)
	assert.True(t, response.Success)

	userOut, ok := response.Content.(*models.User)

	assert.True(t, ok)
	assert.True(t, userOut.Id == 98)
	assert.True(t, userOut.UserName == "Steve Rogers")
	assert.True(t, userOut.Location == "New York, NY")
}

func TestRestPut(t *testing.T) {

	response := models.Response{
		Content: new(models.User),
	}

	userIn := models.User{
		Id:       99,
		UserName: "Natasha Romanov",
		Location: "New York, NY",
	}

	err := wreckerClient.Put("/users").
		WithBody(userIn).
		Into(&response).
		Execute()

	assert.True(t, err == nil)
	assert.True(t, response.Success)

	userOut, ok := response.Content.(*models.User)

	assert.True(t, ok)
	assert.True(t, userOut.Id == 99)
	assert.True(t, userOut.UserName == "Natasha Romanov")
	assert.True(t, userOut.Location == "New York, NY")
}
