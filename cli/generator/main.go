package main

import (
	"fmt"
	"strconv"
	"words_counter/pkg"

	"github.com/spf13/cobra"
)

// to run the cli: go run cli/generator/main.go --size 10m --count 5 --path data/
func main() {
	var size, count, path string

	// create a new cobra command
	cmd := &cobra.Command{
		Use:   "generator",
		Short: "Generate files",
		RunE: func(cmd *cobra.Command, args []string) error {
			// parse the size parameter into bytes
			_, err := pkg.ParseSize(size)
			if err != nil {
				return err
			}

			// convert count from string to integer
			countInt, err := strconv.Atoi(count)
			if err != nil {
				return fmt.Errorf("invalid count: %v", err)
			}

			// generate the files
			err = pkg.GenerateFiles(size, countInt, path)
			if err != nil {
				return fmt.Errorf("failed to generate files: %v", err)
			}

			return nil
		},
	}

	// add size flag to the command
	cmd.Flags().StringVar(&size, "size", "1kb", "Size of the generated file")

	// add count flag to the command
	cmd.Flags().StringVar(&count, "count", "1", "Number of newly generated files")

	// add path flag to the command
	cmd.Flags().StringVar(&path, "path", "data", "Path to store the generated files")

	// mark the count flag as required
	err := cmd.MarkFlagRequired("count")
	if err != nil {
		panic(fmt.Errorf("failed to mark count flag as required: %v", err))
	}

	// execute the command
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
