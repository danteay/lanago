package products

import (
	"os"
	"testing"

	"github.com/danteay/lanago/api/libs/logger"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

var testLogger = new(logger.Logger)

func loadTestServiceEnv() {
	filepath := os.Getenv("CONFIG_FILE")

	err := godotenv.Load(os.ExpandEnv(filepath))
	if err != nil {
		panic(err)
	}
}

func TestProductsServiceStructure(t *testing.T) {
	ps := new(ProductsService)

	assert.Nil(t, ps.Logger)
	assert.EqualValues(t, "", ps.ServiceURL)
}

func TestProductsServiceInit(t *testing.T) {
	loadTestServiceEnv()

	ps := new(ProductsService)
	ps.Init(testLogger)
	ps.Logger.Init()

	assert.EqualValues(t, os.Getenv("PRODUCTS_SERVICE_URL"), ps.ServiceURL)
	assert.NotNil(t, ps.Logger)
}

func TestProductsServiceGetProducts(t *testing.T) {
	loadTestServiceEnv()

	ps := new(ProductsService)
	ps.Init(testLogger)
	ps.Logger.Init()

	_, err := ps.GetProducts()

	assert.NoError(t, err)
}

func TestProductsServiceFindProduct(t *testing.T) {
	loadTestServiceEnv()

	ps := new(ProductsService)
	ps.Init(testLogger)
	ps.Logger.Init()

	prod, err := ps.FindProduct("TSHIRT")

	assert.NoError(t, err)
	assert.EqualValues(t, prod.Code, "TSHIRT")
}
