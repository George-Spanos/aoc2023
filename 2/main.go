package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type LineResult struct {
	Id    int
	Red   int
	Green int
	Blue  int
	Line  string
}
type TurnResult struct {
	Red   int
	Green int
	Blue  int
}

const (
	red   = 12
	green = 13
	blue  = 14
)

func main() {
	challenge2()
}

func challenge1() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatalln("failed to read input file")
	}
	scanner := bufio.NewScanner(file)
	results := make([]LineResult, 0)
	for scanner.Scan() {
		line := scanner.Text()
		r := parseLine(line)
		results = append(results, r)
	}
	fmt.Println(calculateIdsSum(results))
}

func challenge2() {

	file, err := os.Open("input")
	if err != nil {
		log.Fatalln("failed to read input file")
	}
	scanner := bufio.NewScanner(file)
	results := make([]LineResult, 0)
	for scanner.Scan() {
		line := scanner.Text()
		r := parseLine(line)
		results = append(results, r)
	}
	fmt.Println(calculatCubePower(results))
}
func parseLine(line string) LineResult {
	// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	s1 := strings.Split(line, ":")
	sId := strings.Split(s1[0], " ")[1]
	id, err := strconv.Atoi(sId)
	if err != nil {
		log.Fatalf("failed to convert %v to int", sId)
	}
	lineResult := LineResult{
		Id:    id,
		Red:   0,
		Green: 0,
		Blue:  0,
		Line:  line,
	}

	turns := strings.Split(s1[1], ";")
	for _, turn := range turns {
		turnResult := calculateTurnResult(turn)
		if lineResult.Blue < turnResult.Blue {
			lineResult.Blue = turnResult.Blue
		}
		if lineResult.Green < turnResult.Green {
			lineResult.Green = turnResult.Green
		}
		if lineResult.Red < turnResult.Red {
			lineResult.Red = turnResult.Red
		}
	}
	return lineResult
}
func calculateTurnResult(turn string) TurnResult {
	turnResult := TurnResult{
		Red:   0,
		Green: 0,
		Blue:  0,
	}
	colors := strings.Split(turn, ",")
	for i := range colors {
		v := strings.Split(strings.TrimPrefix(colors[i], " "), " ")
		if v[1] == "blue" {
			intValue, err := strconv.Atoi(v[0])
			if err != nil {
				log.Fatalln("failed to convert colors value to number")
			}
			turnResult.Blue = intValue
		} else if v[1] == "green" {
			intValue, err := strconv.Atoi(v[0])
			if err != nil {
				log.Fatalln("failed to convert colors value to number")
			}
			turnResult.Green = intValue
		} else if v[1] == "red" {
			intValue, err := strconv.Atoi(v[0])
			if err != nil {
				log.Fatalln("failed to convert colors value to number")
			}
			turnResult.Red = intValue
		} else {
			log.Fatalln("no color value.")
		}
	}
	return turnResult
}
func calculateIdsSum(results []LineResult) int {
	sum := 0
	for i := range results {
		if results[i].Blue <= blue && results[i].Green <= green && results[i].Red <= red {
			sum += results[i].Id
		}
	}
	return sum
}
func calculatCubePower(results []LineResult) int {
	powerSum := 0
	for i := range results {
		r := results[i]
		powerSum += r.Red * r.Green * r.Blue
	}
	return powerSum
}
