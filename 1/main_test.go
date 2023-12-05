package main

import "testing"

func TestWordIsNumber(t *testing.T) {
	input := map[string]bool{
		"one":   true,
		"asdq":  false,
		"":      false,
		"five":  true,
		"seven": true,
	}
	for k, v := range input {
		_, found := wordIsNumber(k, 0)
		if found != v {
			t.Fatalf("%v is number: %v. Expected %v. Got %v", k, found, v, found)
		}
	}
}
