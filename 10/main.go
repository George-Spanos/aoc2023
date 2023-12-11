package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Position struct {
	x, y int
}

type Pipe struct {
	Position
	symbol string
}

func (p *Pipe) PointsTo() ([]int, []int, error) {
	switch p.symbol {
	case "|":
		return []int{p.y + 1, p.x}, []int{p.y - 1, p.x}, nil
	case "-":
		return []int{p.y, p.x + 1}, []int{p.y, p.x - 1}, nil
	case "F":
		return []int{p.y + 1, p.x}, []int{p.y, p.x + 1}, nil
	case "J":
		return []int{p.y, p.x - 1}, []int{p.y - 1, p.x}, nil
	case "7":
		return []int{p.y, p.x - 1}, []int{p.y + 1, p.x}, nil
	case "L":
		return []int{p.y, p.x + 1}, []int{p.y - 1, p.x}, nil
	}
	return nil, nil, fmt.Errorf("failed to get point %v PointsTo", p)
}
func (p *Pipe) Right(allPipes map[int]map[int]string) Pipe {
	x, y := p.x+1, p.y
	pipeY, found := allPipes[y]
	if !found {
		log.Fatalln("right now found", p)
	}
	symbol, found := pipeY[x]
	if !found {
		log.Fatalln("right now found", p)
	}
	return Pipe{
		Position: Position{
			x: x,
			y: y,
		},
		symbol: symbol,
	}
}
func (p *Pipe) Left(allPipes map[int]map[int]string) Pipe {
	x, y := p.x-1, p.y
	pipeY, found := allPipes[y]
	if !found {
		log.Fatalln("right now found", p)
	}
	symbol, found := pipeY[x]
	if !found {
		log.Fatalln("right now found", p)
	}
	return Pipe{
		Position: Position{
			x: x,
			y: y,
		},
		symbol: symbol,
	}
}
func (p *Pipe) Down(allPipes map[int]map[int]string) Pipe {
	x, y := p.x, p.y+1
	pipeY, found := allPipes[y]
	if !found {
		log.Fatalln("right now found", p)
	}
	symbol, found := pipeY[x]
	if !found {
		log.Fatalln("right now found", p)
	}
	return Pipe{
		Position: Position{
			x: x,
			y: y,
		},
		symbol: symbol,
	}
}
func (p *Pipe) Up(allPipes map[int]map[int]string) Pipe {
	x, y := p.x, p.y-1
	pipeY, found := allPipes[y]
	if !found {
		log.Fatalln("right now found", p)
	}
	symbol, found := pipeY[x]
	if !found {
		log.Fatalln("right now found", p)
	}
	return Pipe{
		Position: Position{
			x: x,
			y: y,
		},
		symbol: symbol,
	}
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
	lines := strings.Split(strings.TrimSpace(string(file)), "\n")
	pipes, startingPosition := getPipes(lines)
	startingPipe := Pipe{
		symbol: pipes[startingPosition.y][startingPosition.x],
		Position: Position{
			x: startingPosition.x,
			y: startingPosition.y,
		},
	}
	adjacentStartingPipes := getAdjacentPipes(pipes, startingPipe)
	i := 1
	prevPipe1 := startingPipe
	pipe1 := adjacentStartingPipes[0]

	prevPipe2 := startingPipe
	pipe2 := adjacentStartingPipes[1]
	// add a soft limit to the endless loop for safety
	for i < 10_000_000 {
		i++
		nextPipe1 := findNextPipe(prevPipe1.Position, pipe1, pipes)
		if nextPipe1 == nil {
			log.Fatalln("there is no next pipe", pipe1)
		}
		prevPipe1, pipe1 = pipe1, *nextPipe1
		if areEqual(pipe1, pipe2) {
			return i
		}
		nextPipe2 := findNextPipe(prevPipe2.Position, pipe2, pipes)
		if nextPipe2 == nil {
			log.Fatalln("there is no next pipe", pipe2)
		}
		prevPipe2, pipe2 = pipe2, *nextPipe2
		if areEqual(pipe1, pipe2) {
			return i
		}
	}
	return -1
}
func part2(file []byte) int {
	return 0
}
func getPipes(lines []string) (map[int]map[int]string, Position) {
	// maps represent y->x->symbol
	pipes := make(map[int]map[int]string)
	startingPipePosition := Position{x: -1, y: -1}
	for i, line := range lines {
		pipes[i] = make(map[int]string)
		for j := range line {
			if isPipeSymbol(line[j]) {
				position := Position{
					x: j,
					y: i,
				}
				pipes[i][j] = string(line[j])
				if string(line[j]) == "S" {
					startingPipePosition = position
				}
			}
		}
	}
	return pipes, startingPipePosition
}
func isPipeSymbol(s byte) bool {
	return string(s) != "."
}
func getAdjacentPipes(allPipes map[int]map[int]string, pipe Pipe) []Pipe {
	adjacentPipes := make([]Pipe, 0)
	directions := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	for _, direction := range directions {
		x, y := pipe.x+direction[1], pipe.y+direction[0]
		pipeY, found := allPipes[y]
		if !found {
			continue
		}
		symbol, found := pipeY[x]
		if !found || !canBeAdjacent(pipe.symbol, symbol, direction) {
			continue
		}
		pipe := Pipe{
			symbol: symbol,
			Position: Position{
				y: y,
				x: x,
			},
		}
		adjacentPipes = append(adjacentPipes, pipe)
	}
	if len(adjacentPipes) > 2 {
		log.Fatalln("adjacent pipes should be always <=2")
	}
	return adjacentPipes
}
func findNextPipe(prevPipePosition Position, currPipe Pipe, allPipes map[int]map[int]string) *Pipe {
	adjPipes := getAdjacentPipes(allPipes, currPipe)
	for _, adjPipe := range adjPipes {
		if adjPipe.x != prevPipePosition.x || adjPipe.y != prevPipePosition.y {
			return &adjPipe
		}
	}
	return nil
}
func areEqual(pipeA, pipeB Pipe) bool {
	return pipeA.symbol == pipeB.symbol && pipeA.x == pipeB.x && pipeA.y == pipeB.y
}

// Given the instructions, not all Symbols can be adjacent.
//
// | is a vertical pipe connecting north and south.
//
// - is a horizontal pipe connecting east and west.
//
// L is a 90-degree bend connecting north and east.
//
// J is a 90-degree bend connecting north and west.
//
// 7 is a 90-degree bend connecting south and west.
//
// F is a 90-degree bend connecting south and east.
func canBeAdjacent(a, b string, direction []int) bool {
	switch a {
	case "|":
		return isUp(direction, b) || isDown(direction, b)
	case "-":
		return isRight(direction, b) || isLeft(direction, b)
	case "7":
		return isDown(direction, b) || isLeft(direction, b)
	case "F":
		return isDown(direction, b) || isRight(direction, b)
	case "L":
		return isUp(direction, b) || isRight(direction, b)
	case "J":
		return isLeft(direction, b) || isUp(direction, b)
	case "S":
		return isLeft(direction, b) || isRight(direction, b) || isDown(direction, b) || isUp(direction, b)
	}
	return false
}
func isRight(direction []int, symbol string) bool {
	return direction[1] == 1 && (symbol == "-" || symbol == "J" || symbol == "7")
}
func isLeft(direction []int, symbol string) bool {
	return direction[1] == -1 && (symbol == "L" || symbol == "F" || symbol == "-")
}
func isUp(direction []int, symbol string) bool {
	return direction[0] == -1 && (symbol == "7" || symbol == "F" || symbol == "|")
}
func isDown(direction []int, symbol string) bool {
	return direction[0] == 1 && (symbol == "L" || symbol == "J" || symbol == "|")
}
