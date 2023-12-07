package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type NumRange struct {
	destStart int
	srcStart  int
	length    int
}

func main() {
	file, err := os.ReadFile("input")
	if err != nil {
		log.Fatalln("failed to read input file")
	}
	fmt.Println("Closest location:", part1(file))
	fmt.Println("???:", part2(file))

}
func part1(file []byte) int {
	lines := strings.Split(string(file), "\n")
	seeds := make([]int, 0)
	ranges := make([][]NumRange, 0)
	var tempRange []NumRange
	for i, line := range lines {
		if len(line) == 0 {
			continue
		}
		if strings.Contains(line, "map") {
			if tempRange != nil {
				ranges = append(ranges, tempRange)
			}
			tempRange = make([]NumRange, 0)
			continue
		}
		if i == 0 {
			seeds = readSeeds(line)
			continue
		}
		r := createRange(line)
		tempRange = append(tempRange, *r)
	}
	ranges = append(ranges, tempRange)
	closestLocation := findClosestLocation(seeds, ranges)
	return closestLocation
}
func part2(file []byte) int {
	lines := strings.Split(string(file), "\n")
	seedRange := readSeedRanges(lines[0])
	ranges := make([][]NumRange, 0)
	var tempRange []NumRange
	for _, line := range lines[1:] {
		if len(line) == 0 {
			continue
		}
		if strings.Contains(line, "map") {
			if tempRange != nil {
				ranges = append(ranges, tempRange)
			}
			tempRange = make([]NumRange, 0)
			continue
		}

		r := createRange(line)
		tempRange = append(tempRange, *r)
	}
	ranges = append(ranges, tempRange)
	closestLocation := findClosestLocationRange(seedRange, ranges)
	return closestLocation
}

func readSeeds(line string) []int {
	s := strings.Split(line, ":")[1]
	seeds := parseStringToIntSlice(s)
	return seeds
}
func readSeedRanges(line string) []NumRange {
	s := strings.Split(line, ":")[1]
	seeds := parseStringToIntSlice(s)
	if len(seeds)%2 != 0 {
		log.Fatalln("seed range length should be even.")
	}
	seedRange := make([]NumRange, 0)
	for i := 0; i < len(seeds); i += 2 {
		r := NumRange{
			srcStart:  seeds[i],
			length:    seeds[i+1],
			destStart: -1,
		}
		seedRange = append(seedRange, r)
	}
	return seedRange
}
func parseStringToIntSlice(line string) []int {
	re := regexp.MustCompile("[0-9]+")
	numsStr := re.FindAllString(line, -1)
	nums := make([]int, len(numsStr))
	for i := range numsStr {
		n, err := strconv.Atoi(numsStr[i])
		if err != nil {
			log.Fatalln("failed to convert string to number")
		}
		nums[i] = n
	}
	return nums
}
func createRange(line string) *NumRange {
	r := NumRange{
		destStart: 0,
		srcStart:  0,
		length:    0,
	}
	nums := parseStringToIntSlice(line)
	r.destStart = nums[0]
	r.srcStart = nums[1]
	r.length = nums[2]
	return &r
}

func findNextValue(src int, r *NumRange) int {
	distance := r.destStart - r.srcStart
	return src + distance
}

// given an input int, find its range inside a NumRange slice
func findRange(src int, ranges []NumRange) *NumRange {
	for _, r := range ranges {
		if inRange(src, &r) {
			return &r
		}
	}
	return nil
}
func findDestRange(dest int, ranges []NumRange) *NumRange {
	for _, r := range ranges {
		if inDestRange(dest, &r) {
			return &r
		}
	}
	return nil
}

// returns true if the give in is inside the NumRange
func inRange(src int, r *NumRange) bool {
	return r.srcStart <= src && src < r.length+r.srcStart
}

// returns true if the give in is inside the NumRange
func inDestRange(dest int, r *NumRange) bool {
	return r.destStart <= dest && dest < r.length+r.destStart
}

func findClosestLocation(seeds []int, ranges [][]NumRange) int {
	closestLocation := math.MaxInt64
	for _, seed := range seeds {
		tempValue := seed
		for _, rl := range ranges {
			r := findRange(tempValue, rl)
			if r != nil {
				tempValue = findNextValue(tempValue, r)
			}
		}
		if closestLocation > tempValue {
			closestLocation = tempValue
		}
	}
	return closestLocation
}
func findClosestLocationRange(seeds []NumRange, ranges [][]NumRange) int {
	limit := 100_000_000
	for i := 0; i < limit; i++ {
		tempValue := i
		for j := len(ranges) - 1; j >= 0; j-- {
			r := findDestRange(tempValue, ranges[j])
			if r != nil {
				tempValue += r.srcStart - r.destStart
			}
		}
		seedRange := findRange(tempValue, seeds)
		if seedRange != nil {
			return i
		}
	}
	log.Fatalln("no seed found")
	return -1
}
