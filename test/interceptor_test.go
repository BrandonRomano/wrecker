package test

import (
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/BrandonRomano/wrecker"
	"github.com/BrandonRomano/wrecker/test/models"
	"github.com/stretchr/testify/assert"
)

func TestInterceptorGet(t *testing.T) {
	// Creating a custom client with interceptor
	wreckerClient := &wrecker.Wrecker{
		BaseURL: "http://localhost:" + os.Getenv("PORT"),
		HttpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		DefaultContentType: "application/x-www-form-urlencoded",
		RequestInterceptor: func(req *wrecker.Request) error {
			req.URLParam("id", "1")
			return nil
		},
	}

	response := models.Response{
		Content: new(models.User),
	}

	httpResponse, err := wreckerClient.Get("/users").
		Into(&response).
		Execute()

	if err != nil {
		t.Error("Error performing Get /users with TestInterceptorGet")
	}

	assert.Equal(t, http.StatusOK, httpResponse.StatusCode)
	assert.Equal(t, response.Content.(*models.User).UserName, "BrandonRomano")
}
