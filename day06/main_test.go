package main

import (
	"testing"
)

func Test_count(t *testing.T) {
	in := []int{0, 2, 7, 0}
	if want, got := 5, count(in); want != got {
		t.Errorf("count() = %d, want %d", got, want)
	}
	if want, got := 5, count(in); want != got {
		t.Errorf("count() = %d, want %d", got, want)
	}
}
