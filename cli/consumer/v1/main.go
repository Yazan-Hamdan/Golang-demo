package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"words_counter/pkg"
	"words_counter/utils"
)

// to run the cli: go run cli/consumer/v1/main.go --path data/
func main() {
	var path string

	rootCmd := &cobra.Command{
		Use:   "consumer",
		Short: "A CLI to count the frequency of each word in all text files in the directory",
		RunE: func(cmd *cobra.Command, args []string) error {
			return printWordFrequencies(path)
		},
	}

	// register command line flags
	rootCmd.Flags().StringVar(&path, "path", "", "The path to the directory containing the text files")

	if err := rootCmd.Execute(); err != nil {
		log.Printf("Error occurred while executing command: %v", err)
		// The exit code of 1 indicates that the program terminated abnormally due to an error
		os.Exit(1)
	}
}

func printWordFrequencies(path string) error {
	// Validate the existence of the directory
	err := utils.ValidateDirectoryExistence(path)
	if err != nil {
		// Log the error
		log.Printf("Error validating directory existence %s: %v", path, err)
		return err
	}

	// Count the frequency of each word in all text files in the directory
	wordFreqs, err := pkg.CountWordsInDir(path)
	if err != nil {
		// Log the error
		log.Printf("Error counting words in directory %s: %v", path, err)
		return err
	}

	// Print the frequency of each word to the console
	for word, freq := range wordFreqs {
		fmt.Printf("%s: %d\n", word, freq)
	}

	return nil
}
