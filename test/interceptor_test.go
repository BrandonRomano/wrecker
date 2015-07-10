package test

import (
	"fmt"
	"github.com/brandonromano/wrecker"
	"os"
	"testing"
)

func TestInterceptorGet(t *testing.T) {

	wrecker.New("http://localhost:" + os.Getenv("PORT"))

	newPlsWork := wrecker.NewPlsWork()
	fmt.Println(newPlsWork.Value)

	plswork := wrecker.PlsWork{
		Value: 100,
	}

	fmt.Println(plswork.Value)

	/*
		response2 := models.Response{
			Content: &models.User{},
		}

		httpResponse, err = w.Get("/users").
			URLParam("id", "1").
			Into(&response2).
			Execute()

		if err != nil {
			t.Error("Error performing GET /users")
		}

		assert.Equal(t, http.StatusOK, httpResponse.StatusCode)
		assert.Equal(t, response2.Content.(*models.User).UserName, "BrandonRomano")
	*/
}

/*
func TestInterceptorPut(t *testing.T) {

	w := wrecker.New("http://localhost:" + os.Getenv("PORT"))

	// This sample interceptor will change the request body to something
	// completely different
	w.Intercept(wrecker.Interceptor{
		WreckerRequest: func(r *wrecker.Request) error {

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
*/
