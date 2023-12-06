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
	score := part1(ReadTestFile(t))
	if score != 13 {
		t.Fatalf("incorrect score sum. expected %v. got %v", 13, score)
	}
}

func TestPart2(t *testing.T) {
	scanned := part2(ReadTestFile(t))
	expected := 30
	if scanned != expected {
		t.Fatalf("incorrect scanned sum. expected %v. got %v", expected, scanned)
	}
}
