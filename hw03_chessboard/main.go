package main

import (
	"fmt"
)

func main() {
	fmt.Println("Please, enter board size")
	var size int
	fmt.Scanf("%d", &size)

	chessboard(size)
}

func chessboard(size int) {
	black := "#"
	white := " "
	finalblack := ""
	finalwhite := ""
	for i := 0; i < size; i++ {
		if i%2 == 0 {
			finalblack += black
			finalwhite += white
		} else {
			finalblack += white
			finalwhite += black
		}
	}
	// fmt.Println(finalblack)
	// fmt.Println(finalwhite)
	for i := 0; i < size; i++ {
		if i%2 == 0 {
			fmt.Println(finalblack)
		} else {
			fmt.Println(finalwhite)
		}
	}
}
