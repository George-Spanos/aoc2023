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
	fmt.Println("Closest location:", part1(file))
	fmt.Println("???:", part2(file))

}
func part1(file []byte) int {
	closestLocation := 0
	lines := strings.Split(string(file), "\n")
	seeds := make([]int, 0)
	maps := make([]map[int]int, 0)
	var tempMap map[int]int
	for i := range lines {
		if len(lines[i]) == 0 {
			maps = append(maps, tempMap)
			tempMap = nil
			continue
		}
		if i == 0 {
			seeds = readSeeds(lines[i])
			continue
		}
		m := createMap(lines[i])
		if tempMap == nil {
			tempMap = make(map[int]int)
		}
		for k, v := range m {
			tempMap[k] = v
		}

	}
	for _, seed := range seeds {
		// 1. get the location for each seed. You have to search the []maps created.
		// 2. the last mapped value should be the location of the seed.
		// 3. keep the minimum found and save it to closestLocation
	}
	return closestLocation
}
func part2(file []byte) int {
	log.Fatalln("not yet implemented")
	return 0
}
func readSeeds(line string) []int {
	s := strings.Split(line, ":")[1]
	seeds := parseStringToIntSlice(s)
	return seeds
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
func createMap(line string) map[int]int {
	m := make(map[int]int)
	nums := parseStringToIntSlice(line)
	for i := 0; i < nums[2]; i++ {
		m[nums[1]+i] = nums[0] + i
	}
	return m
}
