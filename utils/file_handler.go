package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// ReadFile reads the contents of a file given its path.
// It returns the file content as a byte slice, and an error if any.
func ReadFile(filePath string) ([]byte, error) {
	// Open the file with read-only permission.
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the content of the file.
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return content, nil
}

// ValidateDirectoryExistence validate if the given directory exists
func ValidateDirectoryExistence(path string) error {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		log.Printf("directory does not exist: %s", path)
		return err
	}
	if !info.IsDir() {
		err = fmt.Errorf("%s is not a directory", path)
		log.Print(err)
		return err
	}
	return nil
}
