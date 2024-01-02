package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func SetDir(file string) (string, error) {
	base, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("get current folder: %w", err)
	}
	abs, err := filepath.Abs(filepath.Join(base, file))
	if err != nil {
		return "", fmt.Errorf("get absolute path: %w", err)
	}
	if _, err := os.Stat(abs); os.IsNotExist(err) {
		if err := os.MkdirAll(abs, os.ModePerm); err != nil {
			return "", fmt.Errorf("dir not create: %w", err)
		}
	}
	return abs, nil
}
