package main

import "fmt"

type Book struct {
	id     int32
	title  string
	author string
	year   int32
	size   int32
	rate   float32
}

type Columns struct {
	year string
	size string
	rate string
}

func main() {
	// Place your code here.

	firstBook := Book{}
	secondBook := Book{}
	firstBook.create(1, "Harry Potter and the Philosopher's Stone", "J. K. Rowling", 1997, 300, 0.7)
	fmt.Println(firstBook)
	secondBook.create(2, "Harry Potter and the Chamber of Secrets", "J. K. Rowling", 1998, 300, 0.5)
	fmt.Println(secondBook)

	// fmt.Println(firstBook.title)
	// fmt.Println(secondBook.title)

	fmt.Println(firstBook.get())
	fmt.Println(secondBook.get())

	books := []Book{
		{3, "Harry Potter and the Prisoner of Azkaban", "J. K. Rowling", 1999, 300, 0.25},
		{4, "Harry Potter and the Goblet of Fire", "J. K. Rowling", 2000, 300, 1.0},
	}
	fmt.Println(books)

	// fmt.Println(isEqual(firstBook, secondBook))

	fromColumns := Columns{}
	fromColumns.bred("year", firstBook, secondBook)
	fromColumns.bred("size", firstBook, secondBook)
	fromColumns.bred("rate", firstBook, secondBook)
	fromColumns.bred("bred kakoy-to", firstBook, secondBook)
}

// Реализуйте методы для установки и получения полей структуры
func (f *Book) create(id int32, title string, author string, year int32, size int32, rate float32) {
	f.id = id
	f.title = title
	f.author = author
	f.year = year
	f.size = size
	f.rate = rate
}

// Реализуйте методы для установки и получения полей структуры
func (f *Book) get() (int32, string, string, int32, int32, float32) {
	return f.id, f.title, f.author, f.year, f.size, f.rate
}

// func isEqual(x Book, y Book) bool {
// 	if x.year >= y.year {
// 		return true
// 	} else {
// 		return false
// 	}
// }

func (c *Columns) bred(arg string, x Book, y Book) {
	if arg == "year" {
		fmt.Println("year was selected")
		if x.year >= y.year {
			fmt.Println(x.title, arg, "greater then", y.title, arg)
		} else {
			fmt.Println(y.title, arg, "greater then", y.title, arg)
		}

	} else if arg == "size" {
		fmt.Println("size was selected")
		if x.size >= y.size {
			fmt.Println(x.title, arg, "greater then", y.title, arg)
		} else {
			fmt.Println(y.title, arg, "greater then", y.title, arg)
		}
	} else if arg == "rate" {
		fmt.Println("rate was selected")
		if x.rate >= y.rate {
			fmt.Println(x.title, arg, "greater then", y.title, arg)
		} else {
			fmt.Println(y.title, arg, "greater then", y.title, arg)
		}
	} else {
		fmt.Println("Ny eto je bred kakoy-to")
	}

}
