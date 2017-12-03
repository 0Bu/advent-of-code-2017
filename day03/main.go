package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("The Manhattan Distance is %v\n", distance(347991))
	fmt.Printf("The Next value larger than %v is %v", 347991, next(347991))
}

func distance(input int) int {

	type point struct {
		x, y, value int
	}

	top := int(math.Ceil(math.Sqrt(float64(input))))
	if top%2 == 0 {
		top++
	}

	a := point{top, top, top * top}
	b := point{1, a.y, a.value + 1 - a.x}
	c := point{1, 1, b.value + 1 - a.x}
	d := point{a.x, 1, c.value + 1 - a.x}

	var pos point
	if input <= d.value {
		pos = point{d.x, d.value + 1 - input, input}
	} else if input <= c.value {
		pos = point{c.value + 1 - input, c.x, input}
	} else if input <= b.value {
		pos = point{b.x, b.value + 1 - input, input}
	} else if input <= a.value {
		pos = point{a.value + 1 - input, a.x, input}
	} else {
		panic(fmt.Sprintf("value %v out of range[%v-%v]", input, a.value, d.value))
	}

	cv := int(math.Ceil(float64(a.x) / 2))
	center := point{cv, cv, 1}

	return int(math.Abs(float64(pos.x-center.x)) + math.Abs(float64(pos.y-center.y)))
}

func next(input int) (n int) {
	dim := 591
	grid := make([][]int, dim)
	for i := range grid {
		grid[i] = make([]int, dim)
	}
	center := int(math.Ceil(float64(dim / 2)))
	grid[center][center] = 1
	return fill(&grid, center, center+1, 0, input)
}

func fill(grid *[][]int, y, x, dir, till int) int { // dir 0-right, 1-up, 2-left, 3-down
	g := *grid
	sum(grid, y, x)
	if g[y][x] > till {
		return g[y][x]
	}
	if dir == 0 { // from right => to up
		if y-1 >= 0 && g[y-1][x] == 0 { // can going up?
			return fill(grid, y-1, x, 1, till)
		} else if x+1 <= len(g) { // still going right
			return fill(grid, y, x+1, 0, till)
		}
	} else if dir == 1 { // from up => to left
		if x-1 >= 0 && g[y][x-1] == 0 { // can left?
			return fill(grid, y, x-1, 2, till)
		} else if y-1 >= 0 { // still going up
			return fill(grid, y-1, x, 1, till)
		}
	} else if dir == 2 { // from left => to down
		if y+1 < len(g) && g[y+1][x] == 0 {
			return fill(grid, y+1, x, 3, till)
		} else if x-1 >= 0 {
			return fill(grid, y, x-1, 2, till)
		}
	} else if dir == 3 { // from down => to right
		if x+1 < len(g) && g[y][x+1] == 0 {
			return fill(grid, y, x+1, 0, till)
		} else if y+1 < len(g) && g[y+1][x] == 0 {
			return fill(grid, y+1, x, 3, till)
		}
	}
	panic(fmt.Sprintf("i'm stuck at y=%v, x=%v", y, x))
}

func sum(grid *[][]int, y, x int) {
	g := *grid
	if x-1 >= 0 && x+1 < len(g) {
		g[y][x] += g[y][x-1] + g[y][x+1]
	}
	if y+1 < len(g) {
		g[y][x] += g[y+1][x]
		if x-1 >= 0 && x+1 < len(g) {
			g[y][x] += g[y+1][x-1] + g[y+1][x+1]
		}
	}
	if y-1 >= 0 {
		g[y][x] += g[y-1][x]
		if x-1 >= 0 && x+1 < len(g) {
			g[y][x] += g[y-1][x+1] + g[y-1][x-1]
		}
	}
}
