package test

import (
	"fmt"
	"github.com/brandonromano/wrecker"
	"os"
	"testing"
)

func TestInterceptorGet(t *testing.T) {

	wrecker.New("http://localhost:" + os.Getenv("PORT"))

	person := wrecker.NewPerson("wow")
	fmt.Println(person.Value)
}
