package wrecker

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

var wreckerClient Wrecker

func init() {
	baseURL := "https://api.github.com"
	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}
	wreckerClient = Wrecker{
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
