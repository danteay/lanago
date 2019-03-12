package products

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/danteay/lanago/api/libs/logger"
	"github.com/danteay/lanago/api/models"
)

// ProductService is the main structure to handle the list of avalable producs
// and it's actions. It has the ServiceURL and a Logger instance porinter to
// debug errors
type ProductsService struct {
	ServiceURL string
	Logger     *logger.Logger
}

// productServiceList is the structure to capture the response of the
// GetProducts method and parsit to return the corresponding list
type productServiceList struct {
	Products []models.Product `json:"products"`
}

// Init initialize receives the logger instances and obtain the service URL, the
// storage the obtained values into the struture properties
func (ps *ProductsService) Init(l *logger.Logger) {
	ps.ServiceURL = os.Getenv("PRODUCTS_SERVICE_URL")
	ps.Logger = l
}

// GetProducts return a slice of product strutures and eror if some exists in
// the call process.
func (ps *ProductsService) GetProducts() ([]models.Product, error) {
	result, err := http.Get(ps.ServiceURL + "/bins/4bwec")

	if err != nil {
		ps.Logger.Error("ProductService.GetProducts", err.Error())
		return nil, err
	}

	defer result.Body.Close()

	var contents []byte

	contents, err = ioutil.ReadAll(result.Body)

	if err != nil {
		ps.Logger.Error("ProductService.GetProducts", err.Error())
		return nil, err
	}

	var data productServiceList

	err = json.Unmarshal(contents, &data)

	if err != nil {
		ps.Logger.Error("ProductService.GetProducts", err.Error())
		return nil, err
	}

	return data.Products, nil
}

// FindProduct search for a specific product code and return the product related
// data, if the code is not found, return nil and an error
func (ps *ProductsService) FindProduct(code string) (*models.Product, error) {
	list, err := ps.GetProducts()

	if err != nil {
		return nil, err
	}

	for _, v := range list {
		if v.Code == code {
			return &v, nil
		}
	}

	return nil, errors.New("not found product")
}
