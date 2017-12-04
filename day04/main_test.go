package main

import (
	"testing"
)

func Test_duplicatesValidator(t *testing.T) {
	if want, got := 1, duplicatesValidator("aa bb cc dd ee"); want != got {
		t.Errorf("duplicatesValidator() = %d, want %d", got, want)
	}
	if want, got := 0, duplicatesValidator("aa bb cc dd aa"); want != got {
		t.Errorf("duplicatesValidator() = %d, want %d", got, want)
	}
	if want, got := 1, duplicatesValidator("aa bb cc dd aaa"); want != got {
		t.Errorf("duplicatesValidator() = %d, want %d", got, want)
	}
}

func Test_anagramsValidator(t *testing.T) {
	if want, got := 1, anagramsValidator("abcde fghij"); want != got {
		t.Errorf("anagramsValidator() = %d, want %d", got, want)
	}
	if want, got := 0, anagramsValidator("abcde xyz ecdab"); want != got {
		t.Errorf("anagramsValidator() = %d, want %d", got, want)
	}
	if want, got := 1, anagramsValidator("a ab abc abd abf abj"); want != got {
		t.Errorf("anagramsValidator() = %d, want %d", got, want)
	}
	if want, got := 1, anagramsValidator("iiii oiii ooii oooi oooo"); want != got {
		t.Errorf("anagramsValidator() = %d, want %d", got, want)
	}
	if want, got := 0, anagramsValidator("oiii ioii iioi iiio"); want != got {
		t.Errorf("anagramsValidator() = %d, want %d", got, want)
	}
}
