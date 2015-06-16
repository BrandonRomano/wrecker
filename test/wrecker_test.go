package test

import (
	"github.com/brandonromano/wrecker"
	"github.com/brandonromano/wrecker/test/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
	"time"
)

var wreckerClient wrecker.Wrecker

func init() {
	go startServer()

	wreckerClient = wrecker.Wrecker{
		BaseURL: "http://localhost:5000",
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
		t.Error("Error!!! TODO")
	}
	assert.True(t, response.Success, "true is true")
}

func TestFailGet(t *testing.T) {
	response := models.Response{
		Content: new(models.User),
	}

	err := wreckerClient.Get("/users", nil, &response)
	if err != nil {
		t.Error("Error!!")
	}

	assert.True(t, !response.Success, "false is false")
}
