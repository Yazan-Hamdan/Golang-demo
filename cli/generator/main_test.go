package main

import (
	"testing"
	"words_counter/pkg"
)

// to run the test: go test -bench=.
func BenchmarkGenerateFiles(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := pkg.GenerateFiles("20mb", 20, "data")
		if err != nil {
			b.Fatalf("Error running GenerateFiles: %v", err)
		}
	}
}
