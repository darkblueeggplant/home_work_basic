package main

import (
	"fmt"

	"github.com/darkblueeggplant/home_work_basic/hw06_testing/chessboard"
	"github.com/darkblueeggplant/home_work_basic/hw06_testing/struct_comparator"
)

func main() {
	// chessboard
	fmt.Println(chessboard.PrintDesk(5))
	// comparator
	firstBook := struct_comparator.Book{}

	firstBook.SetID(1)
	firstBook.SetTitle("Harry Potter and the Philosopher's Stone")
	firstBook.SetAuthor("J. K. Rowling")
	firstBook.SetYear(1997)
	firstBook.SetSize(300)
	firstBook.SetRate(0.7)
	fmt.Println(firstBook)

	secondBook := struct_comparator.Book{}

	secondBook.SetID(2)
	secondBook.SetTitle("Harry Potter and the Chamber of Secrets")
	secondBook.SetAuthor("J. K. Rowling")
	secondBook.SetYear(1998)
	secondBook.SetSize(300)
	secondBook.SetRate(0.5)
	fmt.Println(secondBook)

	comparatorYear := struct_comparator.NewComparator("year")
	fmt.Println(comparatorYear.IsEqual(firstBook, secondBook))

	comparatorSize := struct_comparator.NewComparator("rate")
	fmt.Println(comparatorSize.IsEqual(firstBook, secondBook))

}
