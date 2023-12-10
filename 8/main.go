package main

import (
	"fmt"
	"log"
	"os"
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
	// i'm going to create a map[string][]string. Keys will be the locations and values will be the left/right location respectively.
	// After I'm gonna move through the given path and juml from node to node till I find ZZZ.
	lines := strings.Split(strings.TrimSpace(string(file)), "\n")
	path := lines[0]
	network := getNetwork(lines[3:])
	fmt.Println(path)
	fmt.Println(network)
	return 0
}
func getNetwork(lines []string) map[string][]string {
	nodes := make(map[string][]string)
	for _, line := range lines {
		lineSplit := strings.Split(line, "=")
		nodePath := strings.TrimSpace(lineSplit[0])
		locationsSplit := strings.Split(lineSplit[1], ",")
		left := strings.TrimSpace(locationsSplit[0])[1:]
		right := strings.TrimSpace(locationsSplit[1])[:3]
		nodes[nodePath] = []string{left, right}
	}
	return nodes
}
func part2(file []byte) int {
	return 0
}
