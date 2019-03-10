package config

import (
	"os"
	"path"
	"runtime"

	"github.com/danteay/lanago/api/libs"
	"github.com/danteay/lanago/api/services/products"
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
func (s *ServiceClients) InitServices() {
	s.Redis = GetRedis()

	s.Logger = new(libs.Logger)
	s.Logger.Init()

	s.ProductsService = new(products.ProductsService)
	s.ProductsService.Init(s.Logger)
}

// LoadConfiguration search for a valid configuration file and load it's content
// as environment variables.
//
// It expect for a preconfigured env var CONFIG_FILE that will need to have the
// absolute path to the configuration file. If this env var is not configured it
// will try to find as a relative path in the same path where the server was
// started
func (s *ServiceClients) LoadConfigurations() {
	var filepath string

	if _, err := os.Stat(os.Getenv("CONFIG_FILE")); os.IsNotExist(err) {
		_, filename, _, _ := runtime.Caller(1)
		filepath = path.Join(path.Dir(filename), "/properties.env")
	} else {
		filepath = os.Getenv("CONFIG_FILE")
	}

	err := godotenv.Load(os.ExpandEnv(filepath))
	if err != nil {
		panic(err)
	}
}
