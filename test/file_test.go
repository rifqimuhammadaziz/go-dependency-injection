package test

import (
	"testing"

	"github.com/rifqimuhammadaziz/go-restful-api/simple"
	"github.com/stretchr/testify/assert"
)

func TestConnection(t *testing.T) {
	connection, cleanup := simple.InitializedConnection("nama_file")
	assert.NotNil(t, connection)

	cleanup()
}
