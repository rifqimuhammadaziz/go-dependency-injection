package test

import (
	"fmt"
	"testing"

	"github.com/rifqimuhammadaziz/go-restful-api/simple"
)

func TestSimpleService(t *testing.T) {
	simpleService, err := simple.InitializedService()
	fmt.Println(err)
	fmt.Println(simpleService)
}
