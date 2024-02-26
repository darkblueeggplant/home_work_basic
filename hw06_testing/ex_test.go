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

func Test(t *testing.T) {
	testCases := []struct {
		desc string
	}{
		{
			desc: "",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

		})
	}
}
