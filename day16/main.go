package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	// "github.com/pkg/profile"
)

func main() {
	// defer profile.Start().Stop()

	input, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}
	moves := strings.Split(string(input), ",")
	p := []rune("abcdefghijklmnop")
	p = order(p, moves)
	fmt.Printf("The order of the programs after dance is '%v'\n", string(p))

	cache := make(map[string][]rune)
	for i := int(1e9) % 42; i > 1; i-- { // optimization hint by @zargony - cache size is 42 with 1e9 iterations
		c := string(p)
		if n, ok := cache[c]; ok {
			p = n
			continue
		}
		p = order(p, moves)
		cache[c] = p
	}
	//  fmt.Printf("cache size after 1e9 iterations is %v \n", len(cache))
	fmt.Printf("The order of the programs after a 1e9 dances is '%v'\n", string(p[:]))
}

func order(p []rune, moves []string) []rune {
	for _, m := range moves {
		switch m[0] {
		case 's':
			i, _ := strconv.Atoi(string(m[1:]))
			s := len(p) - i
			p = append(p[s:], p[:s]...)
		case 'x':
			x := strings.Split(m[1:], "/")
			ia, _ := strconv.Atoi(x[0])
			ib, _ := strconv.Atoi(x[1])
			p[ia], p[ib] = p[ib], p[ia]
		case 'p':
			x := strings.Split(m[1:], "/")
			a, b := rune(x[0][0]), rune(x[1][0])
			ia, ib := 0, 0
			for i, r := range p {
				if r == a {
					ia = i
				}
				if r == b {
					ib = i
				}
			}
			p[ia], p[ib] = p[ib], p[ia]
		}
	}
	return p
}
