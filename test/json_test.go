package test

import (
	"github.com/brandonromano/wrecker/test/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostJSON(t *testing.T) {

	response := models.Response{
		Content: new(models.User),
	}

	userIn := models.User{
		Id:       98,
		UserName: "Steve Rogers",
		Location: "New York, NY",
	}

	_, err := wreckerClient.Post("/users").
		Body(userIn).
		Into(&response).
		Execute()

	t.Log(err)
	assert.True(t, err == nil)
	assert.True(t, response.Success)

	userOut, ok := response.Content.(*models.User)

	assert.True(t, ok)
	assert.True(t, userOut.Id == 98)
	assert.True(t, userOut.UserName == "Steve Rogers")
	assert.True(t, userOut.Location == "New York, NY")
}

func TestPutJSON(t *testing.T) {

	response := models.Response{
		Content: new(models.User),
	}

	userIn := models.User{
		Id:       99,
		UserName: "Natasha Romanov",
		Location: "New York, NY",
	}

	_, err := wreckerClient.Put("/users").
		Body(userIn).
		Into(&response).
		Execute()

	t.Log(err)
	assert.True(t, err == nil)
	assert.True(t, response.Success)

	userOut, ok := response.Content.(*models.User)

	assert.True(t, ok)
	assert.True(t, userOut.Id == 99)
	assert.True(t, userOut.UserName == "Natasha Romanov")
	assert.True(t, userOut.Location == "New York, NY")
}

func TestPutString(t *testing.T) {

	response := models.Response{
		Content: new(string),
	}

	_, err := wreckerClient.Put("/status").
		Body("status code green").
		Into(&response).
		Execute()

	t.Log(err)
	assert.True(t, err == nil)
	assert.True(t, response.Success)

	status, ok := response.Content.(*string)

	assert.True(t, ok)
	assert.True(t, *status == "status code green")
}
