package env

import (
	"os"
)

type Env struct {
	DatabaseUrl     string
	TestDatabaseUrl string
}

func GetEnv() Env {
	return Env{
		DatabaseUrl:     GetRequiredEnv("PG_HOST"),
		TestDatabaseUrl: GetRequiredEnv("TEST_DATABASE_URL"),
	}
}

func GetRequiredEnv(key string) string {
	val, exists := os.LookupEnv(key)
	if !exists {
		panic("Missing required environment variable: " + key)
	}

	return val
}
