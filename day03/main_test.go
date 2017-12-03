package main

import (
	"testing"
)

func Test_distance(t *testing.T) {
	if want, got := 0, distance(1); want != got {
		t.Errorf("distance() = %d, want %d", got, want)
	}
	if want, got := 3, distance(12); want != got {
		t.Errorf("distance() = %d, want %d", got, want)
	}
	if want, got := 2, distance(23); want != got {
		t.Errorf("distance() = %d, want %d", got, want)
	}
	if want, got := 31, distance(1024); want != got {
		t.Errorf("distance() = %d, want %d", got, want)
	}
}

func Test_next(t *testing.T) {
	if want, got := 2, next(1); want != got {
		t.Errorf("next() = %d, want %d", got, want)
	}
	if want, got := 54, next(27); want != got {
		t.Errorf("next() = %d, want %d", got, want)
	}
	if want, got := 351, next(331); want != got {
		t.Errorf("next() = %d, want %d", got, want)
	}
	if want, got := 806, next(748); want != got {
		t.Errorf("next() = %d, want %d", got, want)
	}
}
