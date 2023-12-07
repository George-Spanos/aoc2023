package main

import (
	"os"
	"strings"
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
	expected := 35
	closestLocation := part1(ReadTestFile(t))
	if closestLocation != expected {
		t.Fatalf("closest location incorrect. expected %v. got %v", expected, closestLocation)
	}
}
func TestPart2(t *testing.T) {
	expected := 46
	closestLocation := part2(ReadTestFile(t))
	if closestLocation != expected {
		t.Fatalf("closest location incorrect. expected %v. got %v", expected, closestLocation)
	}
}

func TestReadSeeds(t *testing.T) {
	expected := []int{79, 14, 55, 13}
	file := ReadTestFile(t)
	lines := strings.Split(string(file), "\n")
	seeds := readSeeds(lines[0])
	for i := range seeds {
		if seeds[i] != expected[i] {
			t.Fatalf("seed not equal. expected %v. got %v", expected[i], seeds[i])
		}
	}
}
