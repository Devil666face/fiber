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
	dir := filepath.Dir(abs)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return "", err
		}
	}
	return abs, nil
}
