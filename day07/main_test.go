package main

import (
	"strings"
	"testing"
)

const input = `pbga (66)
xhth (57)
ebii (61)
havc (66)
ktlj (57)
fwft (72) -> ktlj, cntj, xhth
qoyq (66)
padx (45) -> pbga, havc, qoyq
tknk (41) -> ugml, padx, fwft
jptl (61)
ugml (68) -> gyxo, ebii, jptl
gyxo (61)
cntj (57)`

func Test_root(t *testing.T) {
	if want, got := "tknk", root(strings.Split(string(input), "\n")).name; want != got {
		t.Errorf("root() = %v, want %v", got, want)
	}
}

func Test_findBalance(t *testing.T) {
	p := root(strings.Split(string(input), "\n"))
	if want, got := 60, findBalance(p); want != got {
		t.Errorf("findBalance() = %v, want %v", got, want)
	}
}
