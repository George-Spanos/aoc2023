package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

type Galaxy struct {
	x, y int
}

func main() {
	file, err := os.ReadFile("input")
	if err != nil {
		log.Fatalln("failed to read input file")
	}
	fmt.Println("Part 1:", part1(file))
	fmt.Println("Part 2:", part2(file))
}

func part1(file []byte) int {
	// parse lines and save symbols
	lines := strings.Split(strings.TrimSpace(string(file)), "\n")
	galaxies := createGalaxies(lines)
	sum := 0
	pairs := createGalaxyPairs(galaxies)
	for _, pair := range pairs {
		distance := calcGalaxyDistance(pair[0], pair[1])
		sum += distance
	}
	return sum
}

func part2(file []byte) int {
	// parse lines and save symbols
	lines := strings.Split(strings.TrimSpace(string(file)), "\n")
	galaxies := createGalaxiesPart2(lines)
	sum := 0
	pairs := createGalaxyPairs(galaxies)
	for _, pair := range pairs {
		distance := calcGalaxyDistance(pair[0], pair[1])
		sum += distance
	}
	return sum
}
func createGalaxiesPart2(lines []string) []Galaxy {
	galaxies := make([]Galaxy, 0)
	maxColumns := len(lines[0])
	emptyRows := 0
	for i, line := range lines {
		lineGalaxies := getLineGalaxies(line, i+emptyRows)
		if len(lineGalaxies) == 0 {
			emptyRows += int(math.Pow10(6))
		} else {
			galaxies = append(galaxies, lineGalaxies...)
		}
	}
	addedCols := 0
	for j := 0; j < maxColumns+addedCols; j++ {
		if IsEmptyColumn(galaxies, j) {
			moveGalaxiesToRightBy(galaxies, j, int(math.Pow10(6)))
			addedCols += int(math.Pow10(6))
			j += int(math.Pow10(6))
		}
	}
	return galaxies
}
func createGalaxies(lines []string) []Galaxy {
	galaxies := make([]Galaxy, 0)
	maxColumns := len(lines[0])
	emptyRows := 0
	for i, line := range lines {
		lineGalaxies := getLineGalaxies(line, i+emptyRows)
		if len(lineGalaxies) == 0 {
			emptyRows++
		} else {
			galaxies = append(galaxies, lineGalaxies...)
		}
	}
	addedCols := 0
	for j := 0; j < maxColumns+addedCols; j++ {
		if IsEmptyColumn(galaxies, j) {
			moveGalaxiesToRight(galaxies, j)
			addedCols++
			j++
		}
	}
	return galaxies
}

func getLineGalaxies(line string, lineIdx int) []Galaxy {
	galaxies := make([]Galaxy, 0)
	for i := range line {
		if string(line[i]) == "#" {
			galaxies = append(galaxies, Galaxy{x: i, y: lineIdx})
		}
	}
	return galaxies
}
func moveGalaxiesToRight(galaxies []Galaxy, idx int) {
	for i := range galaxies {
		if galaxies[i].x > idx {
			galaxies[i].x++
		}
	}
}
func moveGalaxiesToRightBy(galaxies []Galaxy, idx int, by int) {
	for i := range galaxies {
		if galaxies[i].x > idx {
			galaxies[i].x += by
		}
	}
}
func createGalaxyPairs(galaxies []Galaxy) [][]Galaxy {
	galaxyPairs := make([][]Galaxy, 0)
	for i := range galaxies {
		j := i + 1
		for j < len(galaxies) {
			galaxyPairs = append(galaxyPairs, []Galaxy{galaxies[i], galaxies[j]})
			j++
		}
	}
	return galaxyPairs
}

func calcGalaxyPairsLength(galaxies int) int {
	return factorial(galaxies) / (factorial(2) * (factorial(galaxies - 2)))
}
func factorial(number int) int {

	if number == 1 {
		return 1
	}

	factorialOfNumber := number * factorial(number-1)

	return factorialOfNumber
}
func IsEmptyColumn(galaxies []Galaxy, column int) bool {
	for _, galaxy := range galaxies {
		if galaxy.x == column {
			return false
		}
	}
	return true
}

func calcGalaxyDistance(galaxyA, galaxyB Galaxy) int {
	distance := 0
	pivot, static := getPivotX(galaxyA, galaxyB)
	index := 0
	for pivot+index != static {
		index++
	}
	distance += index
	index = 0
	pivot, static = getPivotY(galaxyA, galaxyB)
	for pivot+index != static {
		index++
	}
	distance += index
	return distance
}

func getPivotX(galaxyA, galaxyB Galaxy) (int, int) {
	if galaxyA.x < galaxyB.x {
		return galaxyA.x, galaxyB.x
	} else {
		return galaxyB.x, galaxyA.x
	}
}

func getPivotY(galaxyA, galaxyB Galaxy) (int, int) {
	if galaxyA.y < galaxyB.y {
		return galaxyA.y, galaxyB.y
	} else {
		return galaxyB.y, galaxyA.y
	}
}
