package test

import (
	"github.com/brandonromano/wrecker"
	"github.com/brandonromano/wrecker/interceptors"
	"github.com/brandonromano/wrecker/test/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestInterceptorGet(t *testing.T) {

	var httpResponse *http.Response
	var err error

	w := wrecker.New("http://localhost:" + os.Getenv("PORT"))

	response := models.Response{
		Content: &models.User{},
	}

	// First, show that the request FAILS without the UserId
	httpResponse, err = w.Get("/users").
		Into(&response).
		Execute()

	if err != nil {
		t.Error("Error performing GET /users")
	}

	assert.Equal(t, http.StatusBadRequest, httpResponse.StatusCode)

	// Next, add the RequestInterceptorFunc, and show that the request
	// is successful with it.
	w.Interceptor(wrecker.Interceptor{
		Request: func(r *wrecker.Request) error {
			r.Param("id", "1")
			return nil
		},
	})

	response2 := models.Response{
		Content: &models.User{},
	}

	httpResponse, err = w.Get("/users").
		Param("id", "1").
		Into(&response2).
		Execute()

	if err != nil {
		t.Error("Error performing GET /users")
	}

	assert.Equal(t, http.StatusOK, httpResponse.StatusCode)
	assert.Equal(t, response2.Content.(*models.User).UserName, "BrandonRomano")
}

func TestAuthorizationGet(t *testing.T) {

	w := wrecker.New("http://localhost:" + os.Getenv("PORT"))

	w.Interceptor(interceptors.Authorization("Bearer ABC123"))

	response := models.Response{
		Content: "",
	}

	// First, show that the request FAILS without the UserId
	httpResponse, err := w.Get("/authorization").
		Into(&response).
		Execute()

	if err != nil {
		t.Error("Error performing GET /authorization")
	}

	t.Log(response.Content.(string))
	assert.Equal(t, http.StatusOK, httpResponse.StatusCode)
	assert.Equal(t, response.Content.(string), "Bearer ABC123")
}

func TestInterceptorPut(t *testing.T) {

	w := wrecker.New("http://localhost:" + os.Getenv("PORT"))

	// This sample interceptor will change the request body to something
	// completely different
	w.Interceptor(wrecker.Interceptor{
		Request: func(r *wrecker.Request) error {

			r.Body(models.User{
				Id:       97,
				UserName: "Bruce Banner",
				Location: "Unknown",
			})

			return nil
		},
	})

	response := models.Response{
		Content: new(models.User),
	}

	// This record *should* be overridden by the Interceptor
	userIn := models.User{
		Id:       99,
		UserName: "Natasha Romanov",
		Location: "New York, NY",
	}

	_, err := w.Put("/users").
		Body(userIn).
		Into(&response).
		Execute()

	assert.True(t, err == nil)
	assert.True(t, response.Success)
	t.Log(err)

	userOut, ok := response.Content.(*models.User)

	assert.True(t, ok)
	assert.True(t, userOut.Id == 97)
	assert.True(t, userOut.UserName == "Bruce Banner")
	assert.True(t, userOut.Location == "Unknown")
}
