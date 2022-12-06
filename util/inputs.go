package util

import (
	"os"
	"strings"
)

func GetInputLines(index string) ([]string, error) {
	input, inputErr := os.ReadFile("puzzle" + index + ".txt")

	if inputErr != nil {
		return nil, inputErr
	}

	return strings.Split(strings.TrimSpace(string(input)), "\n"), nil
}
