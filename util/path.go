package util

import "os"

func IsValidPath(path string) error {
	if _, err := os.Stat(path); err != nil {
		return err
	}
	return nil
}
