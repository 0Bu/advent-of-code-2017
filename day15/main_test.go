package main

import (
	"testing"
)

func Test_judge(t *testing.T) {
	a := generator{65, 16807, 1}
	b := generator{8921, 48271, 1}
	if want, got := 1, judge(3, a, b); want != got {
		t.Errorf("judge() = %v, want %v", got, want)
	}
}

func Test_judgeWithCriteria(t *testing.T) {
	a := generator{65, 16807, 4}
	b := generator{8921, 48271, 8}
	if want, got := 1, judge(1056, a, b); want != got {
		t.Errorf("judge() = %v, want %v", got, want)
	}
}
