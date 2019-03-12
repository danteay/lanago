package main

import (
	"os"
	"testing"

	"github.com/danteay/lanago/api/config"
	"github.com/stretchr/testify/assert"
)

func TestEnvironment(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("Panic error:", r)
		}
	}()

	loader := new(config.ServiceClients)
	loader.LoadConfigurations()

	assert.EqualValues(t, "https://api.myjson.com", os.Getenv("PRODUCTS_SERVICE_URL"))
}
