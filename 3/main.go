package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type EnginePart struct {
	x0     int
	x1     int
	y      int
	digits string
}
type Symbol struct {
	y     int
	x     int
	value string
}

var lines = make([]string, 0)

func (n *EnginePart) Value() int {
	v, err := strconv.Atoi(n.digits)
	if err != nil {
		log.Fatalln("failed to convert number to int")
	}
	return v
}
func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatalln("failed to open input file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	numbers := make([]EnginePart, 0)
	symbols := make([]Symbol, 0)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, strings.TrimSpace(line))
		numbers = append(numbers, extractLineNumbers(line, i)...)
		symbols = append(symbols, extractSymbols(line, i)...)
		i++
	}

	sum := part1(symbols, numbers)
	cogs := part2(symbols, numbers)

	fmt.Printf("Sum of Adjecent numbers : %v \n", sum)
	fmt.Printf("Cogs : %v \n", cogs)
}
func part1(symbols []Symbol, numbers []EnginePart) int {
	sum := 0
	for _, symbol := range symbols {
		n := getAdjecentEngineParts(symbol, numbers)
		for _, v := range n {
			sum += v
		}
	}
	return sum
}
func part2(symbols []Symbol, numbers []EnginePart) int {
	cogs := 0
	for _, symbol := range symbols {
		if symbol.value != "*" {
			continue
		}
		engineParts := getAdjecentEngineParts(symbol, numbers)
		if len(engineParts) == 2 {
			cogs += engineParts[0] * engineParts[1]
		}
	}
	return cogs
}
func extractLineNumbers(line string, lineIdx int) []EnginePart {
	numbers := make([]EnginePart, 0)
	var currentNumber *EnginePart
	for i := range line {
		char := line[i]
		_, err := strconv.Atoi(string(char))
		if err != nil {
			if currentNumber != nil {
				numbers = append(numbers, *currentNumber)
				currentNumber = nil
			}
			continue
		}
		if currentNumber == nil {
			currentNumber = &EnginePart{
				x0:     i,
				x1:     i,
				y:      lineIdx,
				digits: string(char),
			}
		} else {
			currentNumber.x1++
			currentNumber.digits += string(char)
		}
	}
	if currentNumber != nil {
		numbers = append(numbers, *currentNumber)
	}
	return numbers
}
func extractSymbols(line string, lineIdx int) []Symbol {
	symbols := make([]Symbol, 0)
	for i := range line {
		char := string(line[i])
		if !isIntOrDot(char) {
			s := Symbol{
				y:     lineIdx,
				x:     i,
				value: string(line[i]),
			}
			symbols = append(symbols, s)
		}
	}
	return symbols
}
func isIntOrDot(c string) bool {
	if c == "." {
		return true
	}
	_, err := strconv.Atoi(c)
	return err == nil
}
func findAdjecentSymbolNumbers(lineNumbers []EnginePart, pos int) []EnginePart {
	engineNumbers := make([]EnginePart, 0)
	for _, number := range lineNumbers {

		if number.x0 == pos+1 || number.x0 == pos || number.x1 == pos || number.x1 == pos-1 || (number.x0 < pos && number.x1 > pos) {
			engineNumbers = append(engineNumbers, number)
		}
	}
	return engineNumbers
}

// given a set of lines returns the engine parts (slice of int)
func getAdjecentEngineParts(symbol Symbol, numbers []EnginePart) []int {
	engineNumbers := make([]EnginePart, 0)
	partsForLine := findAdjecentSymbolNumbers(getNumbersForLine(numbers, symbol.y), symbol.x)
	engineNumbers = append(engineNumbers, partsForLine...)
	if symbol.y > 0 {
		upperLineParts := findAdjecentSymbolNumbers(getNumbersForLine(numbers, symbol.y-1), symbol.x)
		engineNumbers = append(engineNumbers, upperLineParts...)
	}
	if symbol.y < len(lines)-1 {
		downLineParts := findAdjecentSymbolNumbers(getNumbersForLine(numbers, symbol.y+1), symbol.x)
		engineNumbers = append(engineNumbers, downLineParts...)
	}
	engineParts := make([]int, 0)
	for i := range engineNumbers {
		engineParts = append(engineParts, engineNumbers[i].Value())
	}
	return engineParts
}

func getNumbersForLine(numbers []EnginePart, lineIdx int) []EnginePart {
	nums := make([]EnginePart, 0)
	for i := range numbers {
		if numbers[i].y == lineIdx {
			nums = append(nums, numbers[i])
		}
	}
	return nums
}
