package config

import (
	"strconv"

	"github.com/borowiak-m/dynamodb-crud/untils/env"
)

type Config struct {
	Port       int
	Timeout    int
	Dialect    string
	DatabseURI string
}

func GetConfig() Config {
	return Config{
		Port:       parseEnvToInt("PORT", "8080"),
		Timeout:    parseEnvToInt("TIMEOUT", "30"),
		Dialect:    env.GetEnv("DIALECT", "sqlite3"),
		DatabseURI: env.GetEnv("DATABASE_URI", ":memory:"),
	}
}

func parseEnvToInt(envName, defaultValue string) int {
	num, err := strconv.Atoi(env.GetEnv(envName, defaultValue))

	if err != nil {
		return 0
	}
	return num

}
