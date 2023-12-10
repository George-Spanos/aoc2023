package main

import (
	"log"
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
func ReadTestFile2(t *testing.T) []byte {
	t.Helper()
	file, err := os.ReadFile("test_input_2")
	if err != nil {
		t.Fatal("failed to read test input file")
	}
	return file
}
func ReadTestFile3(t *testing.T) []byte {
	t.Helper()
	file, err := os.ReadFile("test_input_3")
	if err != nil {
		t.Fatal("failed to read test input file")
	}
	return file
}

func TestPart1(t *testing.T) {
	file := ReadTestFile(t)
	expected1 := 2
	got := part1(file)
	if got != expected1 {
		t.Fatalf("part 1 test 1 failed. expected %v. got %v", expected1, got)
	}
	expected2 := 6
	file2 := ReadTestFile2(t)
	got = part1(file2)
	if got != expected2 {
		t.Fatalf("part 1 test 2 failed. expected %v. got %v", expected2, got)
	}
}
func TestPart2(t *testing.T) {
	file := ReadTestFile3(t)
	expected := 6
	got := part2(file)
	if got != expected {
		t.Fatalf("part 2 failed. expected %v. got %v", expected, got)
	}
}
func BenchmarkPart2(b *testing.B) {
	file, err := os.ReadFile("test_input_3")
	for i := 0; i < b.N; i++ {
		if err != nil {
			log.Fatal("failed to read test input file")
		}
		part2(file)
	}
}
