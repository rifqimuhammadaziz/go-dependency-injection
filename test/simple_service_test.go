package test

import (
	"fmt"
	"testing"

	"github.com/rifqimuhammadaziz/go-restful-api/simple"
)

func TestSimpleService(t *testing.T) {
	simpleService := simple.InitializedService()
	fmt.Println(simpleService.SimpleRepository)
}
