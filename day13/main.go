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
	fw := firewalls(strings.Split(string(input), "\n"))
	fmt.Printf("The severity of the whole trip is '%v'\n", severity(fw))
	fmt.Printf("The fewest number of picoseconds is '%v'\n", picoseconds(fw))
}

func severity(fw map[int]int) (s int) {
	for i, d := range fw {
		if i%((d-1)*2) == 0 {
			s += i * d
		}
	}
	return
}

func picoseconds(fw map[int]int) (ps int) {
next:
	for ; ; ps++ {
		for i, d := range fw {
			if (ps+i)%((d-1)*2) == 0 {
				continue next
			}
		}
		return
	}
}

// tl;dr use a map https://www.darkcoding.net/software/go-slice-search-vs-map-lookup/
func firewalls(s []string) (f map[int]int) {
	f = make(map[int]int)
	for _, r := range s {
		di := strings.Split(r, ": ")
		i, _ := strconv.Atoi(di[0])
		d, _ := strconv.Atoi(di[1])
		f[i] = d
	}
	return
}
