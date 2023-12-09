package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.ReadFile("input")
	if err != nil {
		log.Fatalln("failed to read input file")
	}
	fmt.Println("Part 1:", part1(file))
	fmt.Println("Part 2:", part2(file))
}

func part1(file []byte) int { return 0 }
func part2(file []byte) int { return 0 }
