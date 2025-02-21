package config

import (
	"errors"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}
var (
	cfg Config
	once   sync.Once
)

func Get() (*Config) {
	once.Do(func() {
		if err := godotenv.Load(); err != nil {
			log.Fatal("error loading .env file")
		}
	
		err := initVariables()
		if err != nil {
			log.Fatalf("%v is not found in the environment", err)
		}
	})

	return &cfg
}

func initVariables() error {
	port, exists := os.LookupEnv("PORT")
	if !exists {
		return errors.New("PORT")
	}

	cfg = Config{
		Port: port,
	}

	return nil
}