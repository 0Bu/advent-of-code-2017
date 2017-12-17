package main

import (
	"testing"
)

func Test_order(t *testing.T) {
	m := []string{"s1", "x3/4", "pe/b"}
	p := []rune("abcde")
	if want, got := "baedc", string(order(p, m)); want != got {
		t.Errorf("order() = %v, want %v", got, want)
	}
}
