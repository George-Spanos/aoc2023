package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("input")
	if err != nil {
		log.Fatalln("failed to read input file")
	}
	cardPoints := part1(file)
	cardsScanned := part2(file)
	fmt.Println("Card poirts:", cardPoints)
	fmt.Println("Cards scanned:", cardsScanned)
}

func part1(file []byte) int {
	totalScore := 0
	for _, card := range strings.Split(string(file), "\n") {
		if len(card) == 0 {
			continue
		}
		winningNumbers, cardNumbers := parseCard(card)
		totalScore += cardScore(winningNumbers, cardNumbers)
	}
	return totalScore
}
func part2(file []byte) int {
	totalCards := 0
	cardMap := make(map[int]int)
	lines := strings.Split(string(file), "\n")
	for i, card := range lines {
		if len(card) == 0 {
			continue
		}
		cardMap[i]++
		winningNumbers, cardNumbers := parseCard(card)
		cardMatches := cardMatches(winningNumbers, cardNumbers)
		cardsToItterate := cardMap[i]
		for k := 0; k < cardsToItterate; k++ {
			for j := 1; j <= cardMatches; j++ {
				cardMap[i+j]++
			}
		}
	}
	for _, v := range cardMap {
		totalCards += v
	}
	return totalCards
}

// get a card string are return winningNumbers and cardNumbers
func parseCard(card string) ([]int, []int) {
	if len(card) == 0 {
		log.Fatalln("card has no length")
	}
	nums := strings.Split(strings.Split(card, ":")[1], "|")
	re := regexp.MustCompile("[0-9]+")
	winningNumbers := make([]int, 0)
	cardNumbers := make([]int, 0)
	for _, n := range re.FindAllString(nums[0], -1) {
		v, err := strconv.Atoi(n)
		if err != nil {
			log.Fatalln("failed to parse number")
		}
		winningNumbers = append(winningNumbers, v)
	}
	for _, n := range re.FindAllString(nums[1], -1) {
		v, err := strconv.Atoi(n)
		if err != nil {
			log.Fatalln("failed to parse number")
		}
		cardNumbers = append(cardNumbers, v)

	}
	return winningNumbers, cardNumbers
}
func cardMatches(winningNumbers []int, cardNumbers []int) int {
	totalMatches := 0
	for _, cardNum := range cardNumbers {
		if slices.Contains(winningNumbers, cardNum) {
			totalMatches++
		}
	}
	if totalMatches < 0 {
		return 0
	}
	return totalMatches
}
func cardScore(winningNumbers []int, cardNumbers []int) int {
	totalMatches := cardMatches(winningNumbers, cardNumbers)
	return int(math.Pow(2, float64(totalMatches-1)))
}
