package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

const input = `
		467..114..
		...*......
		..35..633.
		......#...
		617*......
		.....+.58.
		..592.....
		......755.
		...$.*....
		.664.598..
`
const sum = 4361
const cogsSum = 467835

var numberResults = []int{
	467, 114, 35, 633, 617, 58, 592, 755, 664, 598,
}
var engineResults = []int{
	467, 35, 633, 617, 592, 755, 664, 598,
}

var symbolResults = []Symbol{
	{
		y: 1,
		x: 3,
	},
	{
		y: 3,
		x: 6,
	},
	{
		y: 4,
		x: 3,
	},
	{
		y: 5,
		x: 5,
	},
	{
		y: 8,
		x: 3,
	},
	{
		y: 8,
		x: 5,
	},
}

func TestNumberExtractFromLine(t *testing.T) {
	_, _, numbers := Setup(t)
	for i := range numbers {
		if numbers[i].Value() != numberResults[i] {
			t.Fatalf("failed to parse number correctly. expected %v. got %v", numberResults[i], numbers[i].Value())
		}
	}
}
func TestSymbolExtractFromLine(t *testing.T) {
	_, symbols, _ := Setup(t)
	for i, symbol := range symbols {
		if symbolResults[i].y != symbol.y || symbolResults[i].x != symbol.x {
			t.Fatalf("failed to parse symbols correctly. expected %v. got %v", symbolResults[i], symbol)
		}
	}
}
func TestAdjecentNumbers(t *testing.T) {
	_, symbols, numbers := Setup(t)
	parts := make([]int, 0)
	for _, symbol := range symbols {
		engineParts := getAdjecentEngineParts(symbol, numbers)
		parts = append(parts, engineParts...)
	}
	for _, part := range parts {
		if !slices.Contains(engineResults, part) {
			t.Fatalf("part %v found not in engine results.", part)
		}
	}
}
func TestEngineResults(t *testing.T) {
	_, symbols, numbers := Setup(t)
	parts := make([]int, 0)

	for _, symbol := range symbols {
		engineParts := getAdjecentEngineParts(symbol, numbers)
		parts = append(parts, engineParts...)
	}

	partsSum := 0
	for i := range parts {
		partsSum += parts[i]
	}
	if partsSum != sum {
		t.Fatalf("sums not equal. expected %v. got %v", sum, partsSum)
	}
}
func TestCogs(t *testing.T) {
	_, symbols, numbers := Setup(t)
	cogs := part2(symbols, numbers)
	if cogs != cogsSum {
		t.Fatalf("cogs sums not equal. expected %v. got %v", cogsSum, cogs)
	}
}
func Setup(t *testing.T) ([]string, []Symbol, []EnginePart) {

	t.Helper()
	scanner := bufio.NewScanner((strings.NewReader(strings.TrimSpace(input))))
	symbols := make([]Symbol, 0)
	numbers := make([]EnginePart, 0)
	i := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		lines = append(lines, line)
		symbols = append(symbols, extractSymbols(line, i)...)
		numbers = append(numbers, extractLineNumbers(line, i)...)
		i++
	}
	return lines, symbols, numbers
}
