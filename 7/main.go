package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	HighCard = iota
	Pair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

var figureCardValues = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	// "J": 11, // part 1
	"T": 10,
	"J": 1, // part 2
}

type Hand struct {
	Cards string
	Type  int
	Bid   int
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
	hands := make([]Hand, 0)
	for _, line := range lines {
		hand := parseLine(line)
		hands = append(hands, *hand)
	}
	sort.Slice(hands, func(i, j int) bool {
		return cpmHand(hands[i], hands[j]) == 1
	})
	hands = sortHands(hands, 0, len(hands)-1)
	sum := 0
	for i, hand := range hands {
		sum += hand.Bid * (len(hands) - i)
	}
	return sum
}
func part2(file []byte) int {
	lines := strings.Split(strings.TrimSpace(string(file)), "\n")
	hands := make([]Hand, 0)
	for _, line := range lines {
		hand := parseLine(line)
		hands = append(hands, *hand)
	}
	sort.Slice(hands, func(i, j int) bool {
		return cpmHand(hands[i], hands[j]) == 1
	})
	hands = sortHands(hands, 0, len(hands)-1)
	sum := 0
	for i, hand := range hands {
		sum += hand.Bid * (len(hands) - i)
	}
	return sum
}
func parseLine(line string) *Hand {
	tmp := strings.Split(line, " ")
	cards := tmp[0]
	strBid := tmp[1]
	bid, err := strconv.Atoi(strBid)
	if err != nil {
		log.Fatalf("failed to convert bid %v to int", strBid)
	}
	hand := Hand{
		Cards: cards,
		Bid:   bid,
	}
	hand.Type = getHandType(hand)
	return &hand
}
func getCardValue(card string) int {
	v, err := strconv.Atoi(card)
	if err == nil {
		return v
	}
	figureValue, found := figureCardValues[card]
	if !found {
		log.Fatalf("failed to parse card %v", card)
	}
	return figureValue
}
func cpmHand(handA, handB Hand) int {
	if handA.Type > handB.Type {
		return 1
	}
	if handA.Type < handB.Type {
		return -1
	}
	for i := 0; i < 5; i++ {
		v1 := getCardValue(string(handA.Cards[i]))
		v2 := getCardValue(string(handB.Cards[i]))
		if v1 > v2 {
			return 1
		}
		if v2 > v1 {
			return -1
		}
	}
	return 0
}

func getHandType(hand Hand) int {
	values := make(map[string]int)
	maxValueKey := ""
	maxValue := 0
	jokers := 0
	for i := 0; i < 5; i++ {
		card := string(hand.Cards[i])
		if string(card) == "J" {
			jokers++
			continue
		}
		if _, found := values[card]; found {
			values[card]++
			if values[card] > maxValue {
				maxValue = values[card]
				maxValueKey = string(card)
			}
		} else {
			values[string(hand.Cards[i])] = 1
			if values[card] > maxValue {
				maxValue = values[card]
				maxValueKey = string(card)
			}
		}
	}
	if jokers > 0 {
		values[maxValueKey] += jokers
	}
	switch len(values) {
	case 5:
		return HighCard
	case 4:
		return Pair
	case 3:
		for _, v := range values {
			if v == 3 {
				return ThreeOfAKind
			}
		}
		return TwoPair
	case 2:
		for _, v := range values {
			if v == 3 || v == 2 {
				return FullHouse
			} else {
				return FourOfAKind
			}
		}
	case 1:
		return FiveOfAKind
	default:
		log.Fatalf("did not find a type for hand %v", hand)
	}
	return -1
}
func sortHands(hands []Hand, hi, low int) []Hand {
	if low < hi {
		var p int
		hands, p = partition(hands, low, hi)
		hands = sortHands(hands, low, p-1)
		hands = sortHands(hands, p+1, hi)
	}
	return hands
}
func partition(arr []Hand, low, high int) ([]Hand, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if cpmHand(arr[j], pivot) == 1 {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}
