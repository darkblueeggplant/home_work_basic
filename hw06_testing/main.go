package main

import (
	"fmt"

	"github.com/darkblueeggplant/home_work_basic/hw06_testing/chessboard"
	"github.com/darkblueeggplant/home_work_basic/hw06_testing/struct_comparator"
)

func main() {
	fmt.Println(chessboard.PrintDesk(5))
	firstBook := struct_comparator.Book{}

	// firstBook.SetID(1)
	// firstBook.SetTitle("Harry Potter and the Philosopher's Stone")
	// firstBook.SetAuthor("J. K. Rowling")
	// firstBook.SetYear(1997)
	// firstBook.SetSize(300)
	// firstBook.SetRate(0.7)
	fmt.Println(firstBook)
}
