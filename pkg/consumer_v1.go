package pkg

import (
	"os"
	"path/filepath"
	"strings"
	"words_counter/utils"
)

// CountWordsInDir counts the frequency of each word in all text files in the directory.
func CountWordsInDir(dirPath string) (map[string]int, error) {
	wordFreqs := make(map[string]int)

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if filepath.Ext(path) != ".txt" {
			return nil
		}

		fileContents, err := utils.ReadFile(path)
		if err != nil {
			return err
		}

		words := SplitWords(string(fileContents))
		IncrementWordFreqs(words, wordFreqs)

		return nil
	})

	if err != nil {
		return nil, err
	}

	return wordFreqs, nil
}

// SplitWords splits a string of text into individual words.
func SplitWords(text string) []string {
	words := strings.Fields(strings.ToLower(text))
	for i := 0; i < len(words); i++ {
		words[i] = strings.Trim(words[i], ",.!?\"';:-()[]{}")
	}
	return words
}

// IncrementWordFreqs increments the frequency of each word in the given map.
func IncrementWordFreqs(words []string, wordFreqs map[string]int) {
	for _, word := range words {
		wordFreqs[word]++
	}
}
