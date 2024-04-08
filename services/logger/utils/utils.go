package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func ReadFile(filename string) ([]byte, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func WriteFile(filename string, data []byte) error {
	err := os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func GetCaller() (string, string) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		return "", ""
	}
	file = filepath.Base(file)
	parts := strings.Split(file, ".")
	filename := parts[0]
	return filename, fmt.Sprintf("%d", line)
}
