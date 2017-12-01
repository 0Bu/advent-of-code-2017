package main

import "testing"

func Test_sequent(t *testing.T) {
	if sum := sequent([]int{1, 1, 2, 2}); sum != 3 {
		t.Errorf("sequent() = %d, want %d", sum, 3)
	}
	if sum := sequent([]int{1, 1, 1, 1}); sum != 4 {
		t.Errorf("sequent() = %d, want %d", sum, 4)
	}
	if sum := sequent([]int{1, 2, 3, 4}); sum != 0 {
		t.Errorf("sequent() = %d, want %d", sum, 0)
	}
	if sum := sequent([]int{9, 1, 2, 1, 2, 1, 2, 9}); sum != 9 {
		t.Errorf("sequent() = %d, want %d", sum, 9)
	}
}

func Test_opposed(t *testing.T) {
	if sum := opposed([]int{1, 2, 1, 2}); sum != 6 {
		t.Errorf("opposed() = %d, want %d", sum, 6)
	}
	if sum := opposed([]int{1, 2, 2, 1}); sum != 0 {
		t.Errorf("opposed() = %d, want %d", sum, 0)
	}
	if sum := opposed([]int{1, 2, 3, 4, 2, 5}); sum != 4 {
		t.Errorf("opposed() = %d, want %d", sum, 4)
	}
	if sum := opposed([]int{1, 2, 1, 3, 1, 4, 1, 5}); sum != 4 {
		t.Errorf("opposed() = %d, want %d", sum, 4)
	}
}
