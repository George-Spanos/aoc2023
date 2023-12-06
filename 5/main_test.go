package main

import (
	"os"
	"testing"
)

func ReadTestFile(t *testing.T) []byte {
	t.Helper()
	file, err := os.ReadFile("test_input")
	if err != nil {
		t.Fatal("failed to read test input")
	}
	return file
}

func TestPart1(t *testing.T) {
	t.Fatal("not implemented")
}
func TestPart2(t *testing.T) {
	t.Fatal("not implemented")
}
