package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type AppConfig struct {
	Debug              bool
	Port               string
	DatastoreProjectId string
	DatastoreHost      string
}

var Conf AppConfig

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func Load() {
	Conf.Port = getEnv("APP_PORT", "8080")
	Conf.Debug = getEnvBool("APP_DEBUG", false)
	Conf.DatastoreProjectId = getEnv("DATASTORE_PROJECT_ID", "test-demo")
	Conf.DatastoreHost = getEnv("DATASTORE_HOST", "http://localhost:8081")
}

func getEnv(key, defaultVal string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = defaultVal
	}
	return value
}

func getEnvBool(key string, defaultVal bool) bool {
	s := getEnv(key, strconv.FormatBool(defaultVal))

	v, err := strconv.ParseBool(s)
	if err != nil {
		return false
	}
	return v
}
