package config

import (
	"fmt"
	"os"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func InitSounDiverseDBConfig() DBConfig {
	return DBConfig{
		Host:     getEnv("SOUNDIVERSE_DB_HOST"),
		Port:     getEnv("SOUNDIVERSE_DB_PORT"),
		User:     getEnv("SOUNDIVERSE_DB_USER"),
		Password: getEnv("SOUNDIVERSE_DB_PASSWORD"),
		Name:     getEnv("SOUNDIVERSE_DB_NAME"),
	}
}

func getEnv(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	panic(fmt.Errorf("Error: unable get os env [%s]", key))
}
