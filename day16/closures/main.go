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

	input, err := ioutil.ReadFile("../input")
	if err != nil {
		panic(err)
	}

	var p [16]rune
	copy(p[:], []rune("abcdefghijklmnop"))
	inst := instructions(p, strings.Split(string(input), ","))
	p = order(inst, p)
	fmt.Printf("The order of the programs after dance is '%v'\n", string(p[:]))

	cache := make(map[string]string)
	for i := 1; i < 1e9; i++ {
		if n, ok := cache[string(p[:])]; ok {
			copy(p[:], []rune(n))
			continue
		}
		k := string(p[:])
		p = order(inst, p)
		cache[k] = string(p[:])
	}
	fmt.Printf("The order of the programs after a billion dances is '%v'\n", string(p[:]))
}

func order(inst []instruction, p [16]rune) [16]rune {
	for _, i := range inst {
		p = i(p)
	}
	return p
}

type instruction func([16]rune) [16]rune

func instructions(p [16]rune, moves []string) (i []instruction) {

	spin := func(move string) func([16]rune) [16]rune {
		i, _ := strconv.Atoi(string(move[1:]))
		s := len(p) - i
		return func(p [16]rune) [16]rune {
			a := append(p[s:], p[:s]...)
			copy(p[:], a[:])
			return p
		}
	}

	exchange := func(move string) func([16]rune) [16]rune {
		t := strings.Split(move[1:], "/")
		ia, _ := strconv.Atoi(t[0])
		ib, _ := strconv.Atoi(t[1])
		return func(p [16]rune) [16]rune {
			p[ia], p[ib] = p[ib], p[ia]
			return p
		}
	}

	partner := func(move string) func([16]rune) [16]rune {
		t := strings.Split(move[1:], "/")
		a, b := rune(t[0][0]), rune(t[1][0])
		find := func(a, b rune, p [16]rune) (ia, ib int) {
			for i, r := range p {
				if r == a {
					ia = i
				}
				if r == b {
					ib = i
				}
			}
			return
		}
		return func(p [16]rune) [16]rune {
			ia, ib := find(a, b, p)
			p[ia], p[ib] = p[ib], p[ia]
			return p
		}
	}

	for _, m := range moves {
		switch m[0] {
		case 's':
			i = append(i, spin(m))
		case 'x':
			i = append(i, exchange(m))
		case 'p':
			i = append(i, partner(m))
		}
	}

	return
}
