package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("input")
	if err != nil {
		log.Fatalln("failed to read input file")
	}
	fmt.Println("Part 1:", part1(file))
	fmt.Println("Part 2:", part2(file))
}

func part1(file []byte) int {
	lines := strings.Split(strings.TrimSpace(string(file)), "\n")
	sum := 0
	for _, line := range lines {
		nums := readLine(line)
		prediction, err := findPrediction(nums)
		if err == nil {
			sum += prediction
		}
	}
	return sum
}

func part2(file []byte) int {
	return -1
}
func readLine(line string) []int {
	re := regexp.MustCompile("\\d+")
	strNums := re.FindAllString(line, -1)
	if len(strNums) == 0 {
		log.Fatalln("failed to extract numbers from line")
	}
	nums := make([]int, len(strNums))
	for i := range strNums {
		n, err := strconv.Atoi(strNums[i])
		if err != nil {
			log.Fatalln("failed to convert string to number")
		}
		nums[i] = n
	}
	return nums
}
func findPrediction(numSeries []int) (int, error) {
	if sliceHasZeros(numSeries) {
		return 0, nil
	}
	nextSlice := calcChildSlice(numSeries)
	if len(nextSlice) == 0 {
		return 0, fmt.Errorf("reached empty slice")
	}
	nextPrediction, err := findPrediction(nextSlice)
	if err != nil {
		return 0, err
	}
	return numSeries[len(numSeries)-1] + nextPrediction, nil
}

func calcChildSlice(s []int) []int {
	childSlice := make([]int, len(s)-1)
	for i := 0; i < len(childSlice); i++ {
		childSlice[i] = s[i+1] - s[i]
	}
	return childSlice
}
func sliceHasZeros(s []int) bool {
	for i := range s {
		if s[i] != 0 {
			return false
		}
	}
	return true
}
