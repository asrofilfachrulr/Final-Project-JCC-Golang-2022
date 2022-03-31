package utils

import (
	"os"
	"strconv"
)

func GetEnvWithFallback(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func StringToIntIgnore(s string) int {
	n, _ := strconv.Atoi(s)

	return n
}
