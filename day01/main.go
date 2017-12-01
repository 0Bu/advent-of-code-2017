package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func main() {
	input, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}
	digits := intArray(string(input))
	fmt.Printf("The solution to the sequent captcha is %d\n", sequent(digits))
	fmt.Printf("The solution to the opposed captcha is %d\n", opposed(digits))
}

func sequent(digits []int) (sum int) {
	if digits[0] == digits[len(digits)-1] {
		sum = digits[0]
	}
	for i := 0; i < len(digits)-1; i++ {
		if digits[i] == digits[i+1] {
			sum += digits[i]
		}
	}
	return
}

func opposed(digits []int) (sum int) {
	a, b := digits[:len(digits)/2], digits[len(digits)/2:]
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			sum += a[i] * 2
		}
	}
	return
}

func intArray(s string) []int {
	a := make([]int, len(s))
	for i, r := range s {
		a[i], _ = strconv.Atoi(string(r))
	}
	return a
}
