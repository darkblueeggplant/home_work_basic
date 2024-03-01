package main

import (
	"testing"

	"github.com/darkblueeggplant/home_work_basic/hw06_testing/chessboard"
)

func TestPrintSomething(t *testing.T) {
	const str, want = "Hello", "Hello"
	got := chessboard.PrintSomething(str)
	if got != want {
		t.Errorf("something wrong: got %v, want %v", got, want)
	}
}

func Test(t *testing.T) {
	testCases := []struct {
		desc      string
		str, want string
	}{
		{
			desc: "chessboard.PrintSomething",
			str:  "Hello",
			want: "Hello",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := chessboard.PrintSomething(tC.str)
			if got != tC.want {
				t.Errorf("something wrong: got %v, want %v", got, tC.want)
			}
		})
	}
}
