package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}
	steps := strings.Split(string(input), ",")
	fmt.Printf("The fewest number of steps is '%v'\n", distance(steps))
	fmt.Printf("The furthest he ever got is '%v'\n", furthest(steps))
}

// https://www.redblobgames.com/grids/hexagons/#distances
func distance(steps []string) int {
	count := func(s string) (i int) {
		for _, r := range steps {
			if s == r {
				i++
			}
		}
		return
	}
	a := count("s") - count("n") + count("sw") - count("ne")
	b := count("se") - count("nw") + count("ne") - count("sw")
	max := []int{
		int(math.Abs(float64(a))),
		int(math.Abs(float64(b))),
		int(math.Abs(float64(-a - b))),
	}
	sort.Ints(max)
	return max[len(max)-1]
}

func furthest(steps []string) int {
	var a, b, c int
	for _, r := range steps {
		switch r {
		case "n":
			a--
		case "s":
			a++
		case "se":
			b++
		case "nw":
			b--
		case "ne":
			a--
			b++
		case "sw":
			a++
			b--
		}
		max := []int{
			c,
			int(math.Abs(float64(a))),
			int(math.Abs(float64(b))),
			int(math.Abs(float64(-a - b))),
		}
		sort.Ints(max)
		c = max[len(max)-1]
	}
	return c
}
