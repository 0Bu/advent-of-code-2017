package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}
	ints := str2int(strings.Split(string(input), "\n"))
	copy := append([]int{}, ints...)
	fmt.Printf("Reached the exit on step %v\n", jump(&copy, 0))
	fmt.Printf("Reached the exit with even stranger jumps on step %v\n", jumpEvenStranger(ints))
}

func jump(ints *[]int, position int) int {
	jumps := (*ints)[position]
	(*ints)[position]++
	if position+jumps < len(*ints) {
		return jump(ints, position+jumps) + 1
	}
	return 1
}

func jumpEvenStranger(ints []int) (steps int) {
	for i := 0; i < len(ints); steps++ {
		jumps := ints[i]
		if jumps < 3 {
			ints[i]++
		} else {
			ints[i]--
		}
		i += jumps
	}
	return
}

func str2int(s []string) (a []int) {
	a = make([]int, len(s))
	for i, r := range s {
		a[i], _ = strconv.Atoi(string(r))
	}
	return
}
