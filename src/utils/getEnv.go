package utils

import (
	"os"
	"strconv"
)

func GetEnvAsString(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(key + " not present in .env")
	}
	return value
}

func GetEnvAsInt(key string) int {
	value := os.Getenv(key)
	if value == "" {
		panic(key + " not present in .env")
	}
	res, _ := strconv.Atoi(value)
	return res
}
