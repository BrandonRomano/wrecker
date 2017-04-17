package test

import (
	"fmt"
	"github.com/BrandonRomano/wrecker"
	"os"
	"testing"
)

func TestTravis(t *testing.T) {
	wrecker.New("http://localhost:" + os.Getenv("PORT"))

	tt := wrecker.NewTravisTester("wow")
	fmt.Println(tt.Value)
}
