package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var nums = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func main() {

	file, err := os.Open("input")
	if err != nil {
		log.Fatalln("Failed to open input file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	sum := 0
	pos := 0
	for scanner.Scan() {
		digits := make([]int, 0)
		text := scanner.Text()
		for pos < len(text) {
			char := string(text[pos])
			n, err := strconv.Atoi(char)
			if err == nil {
				digits = append(digits, n)
			} else {
				n, found := wordIsNumber(text, pos)
				if found {
					digits = append(digits, n)
				}
			}
			pos++

		}
		fmt.Println("line", text)
		fmt.Println("Line digits in order", digits)
		if len(digits) > 0 {
			sum += digits[0]*10 + digits[len(digits)-1]
		}
		fmt.Println("loop sum is", sum)
		pos = 0
	}
	fmt.Println("Sum ", sum)
	if err := scanner.Err(); err != nil {
		fmt.Println("Scanner error")
	}
}

// Check if there's a number written in text in the current line sequence
// This function needs to mutate the position if any word is found
func wordIsNumber(line string, pos int) (int, bool) {
	if len(line) == 0 {
		return -1, false
	}
	text := line[pos:]
	if len(text) == 0 {
		return -1, false
	}
	word := string(text[0])
	i := 0
	found := false
	foundValue := -1
	strNums := make([]string, 0)
	for k := range nums {
		strNums = append(strNums, k)
	}
out:
	for len(strNums) > 0 {
		newPrefixes := make([]string, 0)
		for _, k := range strNums {
			if strings.HasPrefix(k, word) {
				if k == word {
					found = true
					foundValue = nums[k]
					break out
				} else {
					newPrefixes = append(newPrefixes, k)
				}
			}
		}
		strNums = make([]string, len(newPrefixes))
		copy(strNums, newPrefixes)
		i++
		if i >= len(text) {
			break
		}
		word = word + string(text[i])
	}
	return foundValue, found
}
