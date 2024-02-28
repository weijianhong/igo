package utils

import (
	"errors"
	"os"
	"path/filepath"
)

func PathExists(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err == nil {
		if fi.IsDir() {
			return true, nil
		}
		return false, errors.New("存在同名文件")
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func GetCurrentDir() (string, error) {
	path, err := os.Executable()
	if err != nil {
		return "", err
	}

	dir, _ := filepath.Split(path)
	return dir, nil
}
