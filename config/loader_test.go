package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServiceClients(t *testing.T) {
	context := new(ServiceClients)

	assert.Nil(t, context.Redis)
	assert.Nil(t, context.Logger)
	assert.Nil(t, context.ProductsService)
}

func TestServiceClientsInit(t *testing.T) {
	context := new(ServiceClients)
	context.InitServices()

	assert.NotNil(t, context.Redis)
	assert.NotNil(t, context.Logger)
	assert.NotNil(t, context.ProductsService)
}
