package main

import (
	"fmt"
	"testing"
)

var input = []struct {
	line   string
	result LineResult
}{
	{
		line:   "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		result: LineResult{Id: 1, Red: 4, Green: 2, Blue: 6},
	},
	{
		line:   "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		result: LineResult{Id: 2, Red: 1, Green: 3, Blue: 4},
	},
	{
		line:   "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		result: LineResult{Id: 3, Red: 20, Green: 13, Blue: 6},
	},
	{
		line:   "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		result: LineResult{Id: 4, Red: 14, Green: 3, Blue: 15},
	},
	{
		line:   "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
		result: LineResult{Id: 5, Red: 6, Green: 3, Blue: 2},
	},
}

func TestLineResult(t *testing.T) {
	for _, v := range input {
		if err := evaluateResult(parseLine(v.line), v.result); err != nil {
			t.Fatalf("FAILED TEST FOR LINE: \n %v.\n %v", v.line, err.Error())
		}
	}
}
func TestChallengeOne(t *testing.T) {
	lineResults := make([]LineResult, 0)
	for i := range input {
		v := input[i]
		lineResult := parseLine(v.line)
		if err := evaluateResult(lineResult, v.result); err != nil {
			t.Fatalf("FAILED TEST FOR LINE: \n %v.\n %v", v.line, err.Error())
		}
		lineResults = append(lineResults, lineResult)
	}
	sum := calculateIdsSum(lineResults)
	if sum != 8 {
		t.Fatalf("sum id is incorrect. Expected 8. Got %v", sum)
	}
}
func TestChallengeTwo(t *testing.T) {
	lineResults := make([]LineResult, 0)
	for i := range input {
		v := input[i]
		lineResult := parseLine(v.line)
		if err := evaluateResult(lineResult, v.result); err != nil {
			t.Fatalf("FAILED TEST FOR LINE: \n %v.\n %v", v.line, err.Error())
		}
		lineResults = append(lineResults, lineResult)
	}
	cubePower := calculatCubePower(lineResults)
	if cubePower != 2286 {
		t.Fatalf("Incorrect cube power. Expected 2424. Got %v", cubePower)
	}
}
func evaluateResult(got LineResult, expected LineResult) error {
	if got.Blue != expected.Blue {
		return fmt.Errorf("False blue value. Expected %v. Got %v", expected.Blue, got.Blue)
	}
	if got.Red != expected.Red {
		return fmt.Errorf("False red value. Expected %v. Got %v", expected.Red, got.Red)
	}
	if got.Green != expected.Green {
		return fmt.Errorf("False green value. Expected %v. Got %v", expected.Green, got.Green)
	}
	return nil
}
