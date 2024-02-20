package main

import (
	"testing"

	"github.com/darkblueeggplant/home_work_basic/hw06_testing/chessboard"
)

func TestPrintSomething(t *testing.T) {
	const str, want = "Goodbye", "Hello"
	got := chessboard.PrintSomething(str)
	if got != want {
		t.Errorf("something wrong: got %v, want %v", got, want)
	}
}
