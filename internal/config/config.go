package config

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type Config struct {
	Env          string
	Debug        bool
	LogToScreen  bool
	LogFile      string
	ErrorLogFile string
	Port         string
}

func Load() *Config {
	godotenv.Load()

	return &Config{
		Env:          getStringEnv("ENV", "DEVELOPMENT"),
		Debug:        getBoolEnv("DEBUG", false),
		LogToScreen:  getBoolEnv("LOG_TO_SCREEN", true),
		LogFile:      getStringEnv("LOG_FILE", "./log/app.log"),
		ErrorLogFile: getStringEnv("ERROR_LOG_FILE", "./log/error.log"),
		Port:         getStringEnv("PORT", "8000"),
	}
}

func getBoolEnv(key string, defaultValue bool) bool {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		return defaultValue
	}
	return boolValue
}

func getStringEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		if key == "PORT" {
			return ":" + defaultValue
		}
		return defaultValue
	}
	if key == "PORT" {
		return ":" + value
	}
	return value
}
