package main

import (
	"os"
	"testing"
)

func ReadTestFile(t *testing.T) []byte {
	t.Helper()
	file, err := os.ReadFile("test_input")
	if err != nil {
		t.Fatal("failed to read test input file")
	}
	return file
}

func TestPart1(t *testing.T) {
	file := ReadTestFile(t)
	expected := 114
	got := part1(file)
	if got != expected {
		t.Fatalf("part 1 failed. expected %v. got %v", expected, got)
	}
}
func TestPart2(t *testing.T) {
	file := ReadTestFile(t)
	expected := 0
	got := part2(file)
	if got != expected {
		t.Fatalf("part 2 failed. expected %v. got %v", expected, got)
	}
}
