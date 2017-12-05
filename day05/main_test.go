package main

import (
	"testing"
)

func Test_move(t *testing.T) {
	if want, got := 5, jump(&[]int{0, 3, 0, 1, -3}, 0); want != got {
		t.Errorf("jump() = %d, want %d", got, want)
	}
}

func Test_jumpEvenStranger(t *testing.T) {
	if want, got := 10, jumpEvenStranger([]int{0, 3, 0, 1, -3}); want != got {
		t.Errorf("jumpEvenStranger() = %d, want %d", got, want)
	}
}
