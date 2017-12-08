package main

import (
	"strings"
	"testing"
)

const input = `b inc 5 if a > 1
a inc 1 if b < 5
c dec -10 if a >= 1
c inc -20 if c == 10`

func Test_largest(t *testing.T) {
	max, sup := largest(strings.Split(string(input), "\n"))
	if want, got := 1, max; want != got {
		t.Errorf("largest() = %v, want %v", got, want)
	}
	if want, got := 10, sup; want != got {
		t.Errorf("largest() = %v, want %v", got, want)
	}
}
