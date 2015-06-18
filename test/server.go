package test

import (
	"encoding/json"
	"github.com/brandonromano/wrecker/test/models"
	"github.com/julienschmidt/httprouter"
	"net"
	"net/http"
	"os"
	"strconv"
)

func startServer() error {
	router := buildRouter()
	listener, err := net.Listen("tcp", ":"+os.Getenv("PORT"))
	if err != nil {
		return err
	}
	http.Serve(listener, router)
	return nil
}

func buildRouter() *httprouter.Router {
	// Creating a router
	router := httprouter.New()
	router.GET("/users", GetUser)
	router.POST("/users", PostUser)
	router.PUT("/users", PutUser)
	router.DELETE("/users/:id", DeleteUser)
	return router
}

func GetUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	response := new(models.Response).Init()
	defer response.Output(writer)

	id, err := strconv.Atoi(request.FormValue("id"))
	if err != nil {
		response.StatusCode = http.StatusBadRequest
		return
	}

	user := new(models.User).Load(id)
	response.Content = user
}

func PostUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	response := new(models.Response).Init()
	defer response.Output(writer)

	// Special handling for REST/JSON encoding
	if request.Header.Get("Content-Type") == "application/json" {

		user := new(models.User)

		if err := json.NewDecoder(request.Body).Decode(&user); err != nil {
			response.StatusCode = http.StatusBadRequest
			response.Content = err.Error()
			return
		}

		response.Content = user

	} else { // Otherwise, standard handling for Form post

		id, err := strconv.Atoi(request.FormValue("id"))
		userName := request.FormValue("user_name")
		location := request.FormValue("location")
		if err != nil || len(userName) == 0 || len(location) == 0 {
			response.StatusCode = http.StatusBadRequest
			return
		}

		user := models.User{
			Id:       id,
			UserName: userName,
			Location: location,
		}
		response.Content = user
	}
}

func PutUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	response := new(models.Response).Init()
	defer response.Output(writer)

	if request.Header.Get("Content-Type") == "application/json" {

		user := new(models.User)

		if err := json.NewDecoder(request.Body).Decode(&user); err != nil {
			response.StatusCode = http.StatusBadRequest
			response.Content = err.Error()
			return
		}

		response.Content = user

	} else {

		id, err := strconv.Atoi(request.FormValue("id"))
		userName := request.FormValue("user_name")
		location := request.FormValue("location")
		if err != nil || len(userName) == 0 && len(location) == 0 {
			response.StatusCode = http.StatusBadRequest
			return
		}

		user := new(models.User).Load(id)
		if len(userName) > 0 {
			user.UserName = userName
		}
		if len(location) > 0 {
			user.Location = location
		}

		response.Content = user
	}
}

func DeleteUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	response := new(models.Response).Init()
	defer response.Output(writer)

	_, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		response.StatusCode = http.StatusBadRequest
		return
	}
}
