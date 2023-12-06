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
	fmt.Println("Closest location:", part1(file))
	fmt.Println("???:", part2(file))

}
func part1(file []byte) int {
	log.Fatalln("not yet implemented")
	return 0
}
func part2(file []byte) int {
	log.Fatalln("not yet implemented")
	return 0
}
