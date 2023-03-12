package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"words_counter/pkg"
)

var (
	path    string
	perFile bool
	perLine bool
)

// to run the cli: go run cli/consumer/v2/main.go --path data/ --perfile true --perline false
func main() {
	rootCmd := &cobra.Command{
		Use: "consumer",
		RunE: func(cmd *cobra.Command, args []string) error {
			return start()
		},
	}

	rootCmd.PersistentFlags().StringVar(&path, "path", "", "path to directory of files")
	rootCmd.PersistentFlags().BoolVar(&perFile, "perfile", false, "boolean flag")
	rootCmd.PersistentFlags().BoolVar(&perLine, "perline", false, "boolean flag")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func start() error {
	// Call the Process function from pkg package
	return pkg.Process(path, perFile, perLine)
}
