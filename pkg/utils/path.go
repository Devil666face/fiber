package utils

import (
	"os"
	"path/filepath"
)

func SetPath(file string) (string, error) {
	base, err := os.Getwd()
	if err != nil {
		return "", err
	}
	abs, err := filepath.Abs(filepath.Join(base, file))
	if err != nil {
		return "", err
	}
	if _, err := os.Stat(abs); os.IsNotExist(err) {
		if err := os.MkdirAll(abs, os.ModePerm); err != nil {
			return "", err
		}
	}
	return abs, nil
}
