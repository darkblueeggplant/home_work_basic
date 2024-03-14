package main

import (
	"fmt"
	"testing"
)

func TestWordCounter(t *testing.T) {
	got := fmt.Sprint(countWords("мама мыла раму раму мыла мама"))
	want := fmt.Sprint(map[string]int{
		"мама": 2,
		"мыла": 2,
		"раму": 2,
	})

	if got != want {
		t.Errorf("something wrong: got %v, want %v", got, want)
	}
}
