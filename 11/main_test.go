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
func TestNumberOfPairs(t *testing.T) {
	expected := 36
	got := calcGalaxyPairsLength(9)
	if got != expected {
		t.Fatalf("calcGalaxyPairsLength failed. expected %v. got %v", expected, got)
	}
}
func TestFinalPositions(t *testing.T) {
	file := ReadTestFile(t, 0)
	lines := strings.Split(strings.TrimSpace(string(file)), "\n")
	expected := []Galaxy{
		{x: 4, y: 0},
		{x: 9, y: 1},
		{x: 0, y: 2},
		{x: 8, y: 5},
		{x: 1, y: 6},
		{x: 12, y: 7},
		{x: 9, y: 10},
		{x: 0, y: 11},
		{x: 5, y: 11},
	}
	galaxies := createGalaxies(lines)
	for i := range expected {
		if galaxies[i].x != expected[i].x {
			t.Fatalf("galaxies.x not equal. expected %v. got %v", expected[i].x, galaxies[i].x)
		}
		if galaxies[i].y != expected[i].y {
			t.Fatalf("galaxies.y not equal. expected %v. got %v", expected[i].y, galaxies[i].y)
		}
	}
}

func TestCalcDistance(t *testing.T) {
	file := ReadTestFile(t, 0)
	lines := strings.Split(strings.TrimSpace(string(file)), "\n")
	expected := 15
	galaxies := createGalaxies(lines)
	got := calcGalaxyDistance(galaxies[0], galaxies[6])
	if got != expected {
		t.Fatalf("Test Calc Galaxy Distance: case 1. expected %v. got %v", expected, got)
	}
	expected = 17
	got = calcGalaxyDistance(galaxies[2], galaxies[5])
	if got != expected {
		t.Fatalf("Test Calc Galaxy Distance: case 2. expected %v. got %v", expected, got)
	}
	expected = 5
	got = calcGalaxyDistance(galaxies[7], galaxies[8])
	if got != expected {
		t.Fatalf("Test Calc Galaxy Distance: case 3. expected %v. got %v", expected, got)
	}
}
func TestPart1(t *testing.T) {
	file := ReadTestFile(t, 0)
	expected := 374
	got := part1(file)
	if got != expected {
		t.Fatalf("part 1 failed. expected %v. got %v", expected, got)
	}

}
func TestPart2(t *testing.T) {
	file := ReadTestFile(t, 2)
	expected := 8410
	got := part2(file)
	if got != expected {
		t.Fatalf("part 2 failed. expected %v. got %v", expected, got)
	}
}
