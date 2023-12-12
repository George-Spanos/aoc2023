package main

import (
	"fmt"
	"os"
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
	got := part1(file)
	if got != expected {
		t.Fatalf("part 1 failed. expected %v. got %v", expected, got)
	}

}
func TestPart2(t *testing.T) {
	file := ReadTestFile(t, 2)
	expected := 4
	got := part2(file)
	if got != expected {
		t.Fatalf("part 2 failed. expected %v. got %v", expected, got)
	}
}
