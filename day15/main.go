package main

import (
	"fmt"
)

func main() {
	a := generator{277, 16807, 1}
	b := generator{349, 48271, 1}
	fmt.Printf("The judge's final count is '%v'\n", judge(40000000, a, b))
	a = generator{277, 16807, 4}
	b = generator{349, 48271, 8}
	fmt.Printf("The judge's final count is '%v'\n", judge(5000000, a, b))
}

type generator struct {
	value  int
	faktor int
	mult   int
}

func generate(g generator) func() int {
	return func() int {
		for {
			g.value = (g.value * g.faktor) % 2147483647
			if g.value%g.mult == 0 {
				return g.value & 0xffff
			}
		}
	}
}

func judge(times int, ag, bg generator) (count int) {
	for a, b := generate(ag), generate(bg); times > 0; times-- {
		if a() == b() {
			count++
		}
	}
	return
}
