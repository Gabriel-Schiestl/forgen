package utils

import "os"

func FindDir(dir string) bool {
	if _, err := os.Stat(dir); err != nil {
		return false
	}

	return true
}

func CreateDir(dir string) error {
	if err := os.MkdirAll(dir, 0700); err != nil {
		return err
	}

	return nil
}

func CreateFile(path string) error {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC,
		0600)
	if err != nil {
		return err
	}

	defer file.Close()
	return nil
}
