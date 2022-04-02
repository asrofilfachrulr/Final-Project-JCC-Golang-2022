package utils

import (
	"testing"
)

func TestCountDigits(t *testing.T) {
	t.Logf("CountDigits(2145): %v\n", CountDigits(2145))
	t.Logf("CountDigits(214584): %v\n", CountDigits(214584))
	t.Logf("CountDigits(2): %v\n", CountDigits(2))
}
