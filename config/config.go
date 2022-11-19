package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config interface {
	Get(key string) string
}

type configImpl struct{}

func NewConfig(filenames ...string) *configImpl {
	err := godotenv.Load(filenames...)
	if err != nil {
		log.Fatal(err)
	}
	return &configImpl{}
}

func (config *configImpl) Get(key string) string {
	return os.Getenv(key)
}
