package test

import (
	"github.com/brandonromano/test/models"
	"github.com/brandonromano/wrecker"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

// ========== Server ==========

func buildRouter() *httprouter.Router {
	// Creating a router
	router := httprouter.New()
	router.GET("users", GetTest)
	return router
}

func GetTest(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id := request.FormValue("id")
	user := models.User{
		Id:       id,
		UserName: "",
	}
	response := new(models.Response).Init()
	response.Output(writer)
}

// ========= Test ==========

var wreckerClient wrecker.Wrecker

func init() {
	// TODO

	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}
	wreckerClient = wrecker.Wrecker{
		BaseURL:    baseURL,
		HttpClient: httpClient,
	}
}

type GithubUser struct {
	Name string `json:"name"`
}

func TestApi(t *testing.T) {
	user := new(GithubUser)
	err := wreckerClient.Get("/users/BrandonRomano", nil, &user)
	if err != nil {
		t.Error("Wow check it")
	}
	assert.Equal(t, user.Name, "Brandon Romano", "Should be brandonromano")
}
