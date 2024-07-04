package config

import (
	"errors"
	"fmt"
	"strconv"
)

type AppConfig struct {
	Port      string `json:"port"`
	DsnDB     string `json:"dns_db"`
	IsDebugDB bool   `json:"is_debug_db"`
}

func LoadAppConfig() (*AppConfig, error) {
	dbPort := Env("DB_PORT", 5432)
	if _, err := strconv.ParseUint(dbPort, 10, 32); err != nil {
		return nil, errors.New("failed to parse db dbPort")
	}
	appPort := Env("PORT", 8080)
	if _, err := strconv.ParseUint(appPort, 10, 32); err != nil {
		return nil, errors.New("failed to parse app dbPort")
	}

	// Connection URL to connect to Postgres Database
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		Env("DB_HOST"),
		dbPort,
		Env("DB_USER"),
		Env("DB_PASSWORD"),
		Env("DB_NAME"),
	)
	//println("DSN: ", dsn)
	isDebugDB := Env("DB_DEBUG", "false") == "true"

	return &AppConfig{
		DsnDB:     dsn,
		Port:      appPort,
		IsDebugDB: isDebugDB,
	}, nil
}
