package pkg

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"unicode"
	"words_counter/utils"
)

func Process(path string, perFile bool, perLine bool) error {
	// Check if the directory exists
	if err := utils.ValidateDirectoryExistence(path); err != nil {
		return err
	}

	// Check if both flags are false
	if !perFile && !perLine {
		return fmt.Errorf("at least one of --perfile or --perline should be set to true")
	}

	// Process files in the directory concurrently using goroutines
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	wordCounts := sync.Map{}

	for _, file := range files {
		if !file.IsDir() {
			wg.Add(1)
			go processFile(filepath.Join(path, file.Name()), perFile, perLine, &wg, &wordCounts)
		}
	}

	wg.Wait()

	// Print word count table
	fmt.Println("Word count table:")
	fmt.Println("==================")
	fmt.Printf("%-20s %s\n", "Word", "Count")
	fmt.Println("------------------")
	wordCounts.Range(func(key, value interface{}) bool {
		fmt.Printf("%-20s %d\n", key.(string), *value.(*int32))
		return true
	})

	return nil
}

func processFile(filePath string, perFile bool, perLine bool, wg *sync.WaitGroup, wordCounts *sync.Map) {
	defer wg.Done()

	// Read file content
	content, err := utils.ReadFile(filePath)
	if err != nil {
		log.Printf("Error reading file %s: %v", filePath, err)
		return
	}

	if perLine {
		lines := strings.Split(string(content), "\n")
		for _, line := range lines {
			countWords(line, wordCounts)
		}
	} else {
		countWords(string(content), wordCounts)
	}

	if perFile {
		fmt.Println("Word count table for", filePath)
		fmt.Println("=========================")
		fmt.Printf("%-20s %s\n", "Word", "Count")
		fmt.Println("-------------------------")
		wordCounts.Range(func(key, value interface{}) bool {
			fmt.Printf("%-20s %d\n", key, *value.(*int32))
			return true
		})
		fmt.Println("-------------------------")
	}
}

func countWords(content string, wordCounts *sync.Map) {
	words := strings.Fields(content)
	for _, word := range words {
		if !isValidWord(word) {
			continue
		}
		value, _ := wordCounts.LoadOrStore(word, new(int32))
		atomic.AddInt32(value.(*int32), 1)
	}
}

func isValidWord(word string) bool {
	if len(word) == 0 {
		return false
	}
	for _, char := range word {
		if !unicode.IsLetter(char) {
			return false
		}
	}
	return true
}
