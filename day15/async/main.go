package main

import (
	"fmt"
)

func main() {
	a := generator{277, 16807, 1, make(chan int)}
	b := generator{349, 48271, 1, make(chan int)}
	fmt.Printf("The judge's final count is '%v'\n", judge(40000000, a, b))
	a = generator{277, 16807, 4, make(chan int)}
	b = generator{349, 48271, 8, make(chan int)}
	fmt.Printf("The judge's final count is '%v'\n", judge(5000000, a, b))
}

type generator struct {
	value  int
	faktor int
	mult   int
	ch     chan int
}

func generate(g generator) {
	for {
		g.value = (g.value * g.faktor) % 2147483647
		if g.value%g.mult == 0 {
			g.ch <- g.value & 0xffff
		}
	}
}

func judge(times int, a, b generator) (count int) {
	go generate(a)
	go generate(b)
	for ; times > 0; times-- {
		if <-a.ch == <-b.ch {
			count++
		}
	}
	return
}
