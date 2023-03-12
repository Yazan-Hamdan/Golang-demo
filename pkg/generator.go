package pkg

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"words_counter/utils"
)

// GenerateFiles generates new files with the given size and count in the given path
func GenerateFiles(size string, count int, path string) error {
	// parse the size parameter into bytes
	sizeBytes, err := ParseSize(size)
	if err != nil {
		return err
	}

	// read the static text to use as a template for the generated files
	staticText, err := utils.ReadFile("resources/static_text.txt")
	if err != nil {
		return fmt.Errorf("failed to read static text: %v", err)
	}

	// create the directory to store the generated files
	err = os.MkdirAll(path, 0755)
	if err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}

	// generate the new files
	for i := 1; i <= count; i++ {
		// create the new file
		filename := fmt.Sprintf("file_%d.txt", i)
		filepath := filepath.Join(path, filename)
		file, err := os.Create(filepath)
		if err != nil {
			return fmt.Errorf("failed to create file: %v", err)
		}
		defer file.Close()

		// write the static text to the file until the desired size is reached
		bytesWritten := 0
		for bytesWritten < sizeBytes {
			bytesToWrite := sizeBytes - bytesWritten
			if bytesToWrite > len(staticText) {
				bytesToWrite = len(staticText)
			}
			_, err := file.Write(staticText[:bytesToWrite])
			if err != nil {
				return fmt.Errorf("failed to write to file: %v", err)
			}
			bytesWritten += bytesToWrite
		}
	}

	return nil
}

// ParseSize parses a string that represents a size in kilobytes, megabytes or gigabytes.
// Returns the size in bytes or an error if the string is invalid.
func ParseSize(size string) (int, error) {
	size = strings.ToLower(size)
	if strings.HasSuffix(size, "b") {
		size = size[:len(size)-1]
	}
	unit := size[len(size)-1:]
	value, err := strconv.Atoi(size[:len(size)-1])
	if err != nil {
		return 0, fmt.Errorf("invalid size value: %s", size[:len(size)-1])
	}

	switch unit {
	case "k":
		return value * 1024, nil
	case "m":
		return value * 1024 * 1024, nil
	case "g":
		return value * 1024 * 1024 * 1024, nil
	default:
		return 0, fmt.Errorf("unrecognized size unit: %s", unit)
	}
}
