package main

import (
	"testing"

	"words_counter/pkg"
)

// to run the test: go test -bench=.
func BenchmarkProcess(b *testing.B) {
	path := "/home/yazan/Desktop/golang/words_counter/data"
	perFile := true
	perLine := false

	for i := 0; i < b.N; i++ {
		perFile = !perFile
		perLine = !perLine
		if err := pkg.Process(path, perFile, perLine); err != nil {
			b.Fatalf("error processing files: %v", err)
		}
	}
}
