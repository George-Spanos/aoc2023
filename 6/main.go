package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	file, err := os.ReadFile("input")
	if err != nil {
		log.Fatalln("failed to read input file")
	}
	fmt.Println("Part 1:", part1(file))
	fmt.Println("Part 2:", part2(file))
}
func parseInput(lines []string) [][]int {
	races := make([][]int, 0)
	if len(lines) != 2 {
		log.Fatalln("lines length is not equal to 2")
	}
	time := strings.Split(lines[0], ":")[1]
	distance := strings.Split(lines[1], ":")[1]
	re := regexp.MustCompile("[0-9]+")
	for _, v := range re.FindAllString(time, -1) {
		num, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalln("failed to convert string to number", v)
		}
		nSlice := []int{num, 0}
		races = append(races, nSlice)
	}
	for i, v := range re.FindAllString(distance, -1) {
		num, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalln("failed to convert string to number", v)
		}
		races[i][1] = num
	}
	return races
}
func part1(file []byte) int {
	lines := strings.Split(strings.TrimSpace(string(file)), "\n")
	races := parseInput(lines)
	product := 1
	for i := range races {
		waysToWin := getWaysToWin(races[i][0], races[i][1])
		product *= waysToWin
	}
	return product
}
func part2(file []byte) int {
	lines := strings.Split(strings.TrimSpace(string(file)), "\n")
	t := strings.Split(lines[0], ":")[1]
	re := regexp.MustCompile("\\d+")
	num := re.FindAllString(t, -1)
	strTime := strings.Join(num, "")
	time, _ := strconv.Atoi(strTime)
	d := strings.Split(lines[1], ":")[1]
	num = re.FindAllString(d, -1)
	strDistance := strings.Join(num, "")
	distance, _ := strconv.Atoi(strDistance)
	waysToWin := getWaysToWin(time, distance)
	return waysToWin
}
func getWaysToWin(time, distance int) int {
	waysToWin := 0
	median := time / 2
	d := measureDistance(median, time-median)
	if d > distance {
		waysToWin++
	}
	wg.Add(2)
	go func() {
		for i := median - 1; i > 0; i-- {
			d := measureDistance(i, time-i)
			if d > distance {
				mu.Lock()
				waysToWin++
				mu.Unlock()
			}
		}
		wg.Done()
	}()
	go func() {
		for i := median + 1; i < time; i++ {
			d := measureDistance(i, time-i)
			if d > distance {
				mu.Lock()
				waysToWin++
				mu.Unlock()
			}
		}
		wg.Done()
	}()
	wg.Wait()
	return waysToWin
}

func measureDistance(velocity int, time int) int {
	return velocity * time
}
