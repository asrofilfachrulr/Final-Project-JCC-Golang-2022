package helper

import (
	"encoding/json"
	"fmt"
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

func StringToUint(s string) (uint, error) {
	n, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(n), nil
}

func CountDigits(n int) int {
	count := 1
	for r := n / 10; r != 0; r /= 10 {
		count += 1
	}
	return count
}

func ValidateDigitInt(n int, min int, max int, desc string) error {
	digits := CountDigits(n)
	if digits < min || digits > max {
		return fmt.Errorf("%s is not valid, digit is too long or too short", desc)
	}
	return nil
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
