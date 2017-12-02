package main

import (
	"testing"
)

func Test_checksum(t *testing.T) {
	str := "5\t1\t9\t5\n7\t5\t3\n2\t4\t6\t8"
	if cm := checksum(str); cm != 18 {
		t.Errorf("checksum() = %d, want %d", cm, 18)
	}
}

func Test_divsum(t *testing.T) {
	str := "5\t9\t2\t8\n9\t4\t7\t3\n3\t8\t6\t5"
	if sum := divsum(str); sum != 9 {
		t.Errorf("divsum() = %d, want %d", sum, 9)
	}
}
