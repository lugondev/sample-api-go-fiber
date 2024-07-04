package config

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvFile(envFile string) {
	if envFile != "" {
		_, err := os.Stat(envFile)
		if err != nil {
			if errors.Is(err, os.ErrNotExist) && envFile != GetDefaultEnvPath() {
				envFile = GetDefaultEnvPath()
				println(envFile + " does not exist, using default config file")
			}
		}
	}
	if envFile != "" {
		_ = os.Setenv("CONFIG_FILE", envFile)
		println("Using config file:", envFile)
	}

}

func GetDefaultEnvPath() string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return wd + "/.env"
}

// Env func to get env value
func Env(key string, df ...interface{}) string {
	defaultValue := ""
	if len(df) > 0 {
		defaultValue = fmt.Sprint(df[0])
	}
	// load .env file
	cfgFile := os.Getenv("CONFIG_FILE")
	if cfgFile != "" {
		err := godotenv.Load(cfgFile)
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
	} else {
		err := godotenv.Load(".env")
		if err != nil {
			_ = godotenv.Load("../../.env")
		}
	}

	val := os.Getenv(key)
	if val == "" {
		return defaultValue
	}
	return val
}
