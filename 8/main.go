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
	lines := strings.Split(strings.TrimSpace(string(file)), "\n")
	path := lines[0]
	network := getNetwork(lines[2:])
	currentNode := "AAA"
	jumps := 0
	found := false
	for !found {
		for _, c := range path {
			jumps++
			currentNode = nextValue(currentNode, string(c), network)
			if currentNode == "ZZZ" {
				found = true
				break
			}
		}
	}
	return jumps
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
	lines := strings.Split(strings.TrimSpace(string(file)), "\n")
	path := lines[0]
	network := getNetwork(lines[2:])
	nodes := getStartingNodes(network)
	jumps := 0
out:
	for {
		for _, instruction := range path {
			for i := range nodes {
				nodes[i] = nextValue(nodes[i], string(instruction), network)
			}
			jumps++
			if nodesOnZ(nodes) {
				break out
			}
			fmt.Println("jumps:", jumps)
		}
	}
	return jumps
}
func getStartingNodes(network map[string][]string) []string {
	startingNodes := make([]string, 0)
	for k := range network {
		if string(k[2]) == "A" {
			startingNodes = append(startingNodes, k)
		}
	}
	return startingNodes
}
func nodesOnZ(nodes []string) bool {
	for _, node := range nodes {
		if string(node[2]) != "Z" {
			return false
		}
	}
	return true
}
func nextValue(node string, instruction string, network map[string][]string) string {
	if instruction == "L" {
		return network[node][0]
	} else if instruction == "R" {
		return network[node][1]
	} else {
		log.Fatalln("failed to parse instruction")
	}
	return ""
}
