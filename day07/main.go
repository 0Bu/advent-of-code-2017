package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type program struct {
	name     string
	weight   int
	children []program
}

// implements sort.Interface https://goo.gl/eT3RPK
type byWeight []program

func (a byWeight) Len() int           { return len(a) }
func (a byWeight) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byWeight) Less(i, j int) bool { return a[i].entireWeight() < a[j].entireWeight() }

// returns entire weight incl. sum of children weight
func (p program) entireWeight() (sum int) {
	if len(p.children) > 0 {
		for i := range p.children {
			sum += p.children[i].entireWeight()
		}
		return p.weight + sum
	}
	return p.weight
}

// returns new program by string with children
func create(s string) (p program) {
	re := regexp.MustCompile(`\w+`)
	t := re.FindAllStringSubmatch(s, -1)
	p.name = t[0][0]
	p.weight, _ = strconv.Atoi(string(t[1][0]))
	for _, r := range t[2:] {
		p.children = append(p.children, program{r[0], 0, nil})
	}
	return
}

// searches for children and parents and order them together
func rebalance(p program, m map[string]program) {
	// seach for children, adopt them and remove them from map
	for i, c := range p.children {
		v, ok := m[c.name]
		if ok {
			p.children[i] = v
			delete(m, c.name)
		}
	}
	// search for parent & get adopted
	adopted := false
	for _, v := range m {
		if getAdopted(p, v) {
			adopted = true
			break
		}
	}
	// no parent found, insert into map
	if !adopted {
		m[p.name] = p
	}
}

// returns true, if a program has found parent and get adopted
func getAdopted(p program, papa program) (adopted bool) {
	for i, c := range papa.children {
		if len(c.children) > 0 {
			if getAdopted(p, c) {
				return true
			}
		}
		if c.name == p.name {
			papa.children[i] = p
			return true
		}
	}
	return
}

// returns root program
func root(lines []string) program {
	m := make(map[string]program)
	for _, line := range lines {
		p := create(line)
		rebalance(p, m)
	}
	for _, v := range m {
		return v
	}
	panic("no parent found")
}

func findBalance(p program) int {
	if c := p.children; len(c) > 0 {
		for i := range c {
			if s := findBalance(c[i]); s > 0 {
				return s
			}
		}
		sort.Sort(byWeight(c))
		if c[0].entireWeight() != c[len(c)-1].entireWeight() {
			balance := c[len(c)-1].entireWeight() - c[0].entireWeight()
			return c[len(c)-1].weight - balance
		}
	}
	return 0
}

func main() {
	input, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}
	p := root(strings.Split(string(input), "\n"))
	fmt.Printf("The name of the bottom program is '%v'\n", p.name)
	fmt.Printf("The weight need to be to balanced is %v\n", findBalance(p))
}
