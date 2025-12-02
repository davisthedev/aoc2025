package utils

import (
	"os"
	"strings"
)

func ReadInputFile(path string) ([]string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return strings.Split(strings.TrimSpace(string(content)), "\n"), nil
}

func ReadInputCsvFile(path string) ([]string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return strings.Split(strings.TrimSpace(string(content)), ","), nil
}
