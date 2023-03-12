package main

import (
	"testing"
)

// to run the test: go test -bench=.
func BenchmarkPrintWordFrequencies(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := printWordFrequencies("/home/yazan/Desktop/golang/words_counter/data")
		if err != nil {
			b.Fatalf("Error running printWordFrequencies: %v", err)
		}
	}
}
