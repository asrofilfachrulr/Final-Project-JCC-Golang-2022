package utils

import (
	"encoding/json"
	"io"
	"io/ioutil"
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

func ParseFromJSON(req io.ReadCloser, data interface{}) error {
	jsonByte, err := ioutil.ReadAll(req)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonByte, data)
	if err != nil {
		return err
	}
	return nil
}
