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
}

func (f *Book) create(id int32, title string, author string, year int32, size int32, rate float32) {
	f.id = id
	f.title = title
	f.author = author
	f.year = year
	f.size = size
	f.rate = rate
}

func (f *Book) get() (int32, string, string, int32, int32, float32) {
	return f.id, f.title, f.author, f.year, f.size, f.rate
}
