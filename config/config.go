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

var (
	APIConf   AppConfig
	CacheConf AppConfig
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func Load() {
	APIConf.Port = getEnv("API_PORT", "8080")
	APIConf.Debug = getEnvBool("API_DEBUG", false)
	APIConf.DatastoreProjectId = getEnv("DATASTORE_PROJECT_ID", "test-demo")
	APIConf.DatastoreHost = getEnv("DATASTORE_HOST", "http://localhost:8081")

	CacheConf.Port = getEnv("CACHE_PORT", "8088")
	CacheConf.Debug = getEnvBool("CACHE_DEBUG", false)
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
