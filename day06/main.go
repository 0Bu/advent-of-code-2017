package main

import (
	"fmt"
)

func main() {
	input := []int{0, 5, 10, 0, 11, 14, 13, 4, 11, 8, 8, 7, 1, 4, 12, 11}
	fmt.Printf("Redistribution cycles passed %v\n", count(input))
	fmt.Printf("2-nd redistribution cycles passed %v\n", count(input)-1) // -1 => don't count output
}

func count(blocks []int) (cycles int) {
	for seen := dejavu(); ; cycles++ {
		if seen(spread(blocks)) {
			return cycles + 1
		}
	}
}

/*
	the clousure is as a Set of seen blocks
	returns true, if it has already a block
*/
func dejavu() func(block string) bool {
	blocks := []string{}
	return func(block string) bool {
		for _, r := range blocks {
			if r == block {
				return true
			}
		}
		blocks = append(blocks, block)
		return false
	}
}

func spread(ints []int) (spreaded string) {
	max, index := maxIndex(ints)
	ints[index] = 0
	for i := index + 1; max > 0; i++ { // start spreading with next array element
		ints[i%len(ints)]++
		max--
	}
	return fmt.Sprint(ints)
}

func maxIndex(ints []int) (max, index int) {
	max = ints[0]
	for i, r := range ints {
		if r > max {
			max = r
			index = i
		}
	}
	return
}
