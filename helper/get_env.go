package helper

import (
	"os"
)

func GetEnv(key string) string {
	value := os.Getenv(key)
	return value
}
