package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func ReadTestFile(t *testing.T, i int) []byte {
	t.Helper()
	filename := "test_input"
	if i > 0 {
		filename += fmt.Sprintf("_%v", i)
	}
	file, err := os.ReadFile(filename)
	if err != nil {
		t.Fatal("failed to read test input file")
	}
	return file
}
func TestPart1(t *testing.T) {
	file := ReadTestFile(t, 0)
	fmt.Println(string(file))
	expected := 4
	lines := strings.Split(string(file), "\n")
	got := part1(lines)
	if got != expected {
		t.Fatalf("part 1 test 1 failed. expected %v. got %v", expected, got)
	}
	file = ReadTestFile(t, 1)
	expected = 8
	got = part1(strings.Split(string(file), "\n"))
	if got != expected {
		t.Fatalf("part 1 test 2 failed. expected %v. got %v", expected, got)
	}
}
func TestPart2(t *testing.T) {
	file := ReadTestFile(t, 2)
	expected := 4
	got := part2(strings.Split(string(file), "\n"))
	if got != expected {
		t.Fatalf("part 2 test 1 failed. expected %v. got %v", expected, got)
	}
	// -----
	file = ReadTestFile(t, 3)
	expected = 8
	got = part2(strings.Split(string(file), "\n"))
	if got != expected {
		t.Fatalf("part 2 test 2 failed. expected %v. got %v", expected, got)
	}
	// -----

	file = ReadTestFile(t, 4)
	expected = 10
	got = part2(strings.Split(string(file), "\n"))
	if got != expected {
		t.Fatalf("part 2 test 3 failed. expected %v. got %v", expected, got)
	}
}
