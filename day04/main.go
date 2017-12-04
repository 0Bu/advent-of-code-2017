package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type validator func(string) int

func main() {
	input, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(input), "\n")
	fmt.Printf("Passwords without duplicates = %v\n", countValidPass(lines, duplicatesValidator))
	fmt.Printf("Passwords without anagrams = %v\n", countValidPass(lines, anagramsValidator))
}

func countValidPass(passwords []string, fn validator) (sum int) {
	for _, pass := range passwords {
		sum += fn(pass)
	}
	return
}

func duplicatesValidator(password string) int {
	// regex: go doesn't support backreferences, eg. (\w+).*\1
	words := strings.Split(password, " ")
	for _, word := range words {
		count := 0
		for _, w := range words {
			if w == word {
				count++
			}
		}
		if count > 1 {
			return 0
		}
	}
	return 1
}

func anagramsValidator(password string) int {
	words := strings.Split(password, " ")
	var sorted []string
	for _, r := range words {
		chars := strings.Split(r, "")
		sort.Strings(chars)
		sorted = append(sorted, strings.Join(chars, ""))
	}
	return duplicatesValidator(strings.Join(sorted, " "))
}
