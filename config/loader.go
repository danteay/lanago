package config

import (
	"lanago/api/libs"
	"lanago/api/services/products"
	"os"

	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
)

// ServiceClients is the struct that has all the needed clients to
// perform operations
type ServiceClients struct {
	Redis           *redis.Client
	Logger          *libs.Logger
	ProductsService *products.ProductsService
}

// Init charge all the env variables and cofigures the ServiceClients
// struct
func (s *ServiceClients) Init() {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	err = godotenv.Load(os.ExpandEnv(path + "/properties.env"))
	if err != nil {
		panic(err)
	}

	s.Redis = GetRedis()

	s.Logger = new(libs.Logger)
	s.Logger.Init()

	s.ProductsService = new(products.ProductsService)
	s.ProductsService.Init(s.Logger)
}
