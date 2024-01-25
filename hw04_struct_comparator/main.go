package main

import (
	"fmt"
)

type Book struct {
	id     int32
	title  string
	author string
	year   int32
	size   int32
	rate   float32
}

type Fields struct {
	byYear int
	bySize int
	byRate int
}

// Реализуйте методы для установки и получения полей структуры.

func (f *Book) setID(id int32) {
	f.id = id
}

func (f *Book) setTitle(title string) {
	f.title = title
}

func (f *Book) setAuthor(author string) {
	f.author = author
}

func (f *Book) setYear(year int32) {
	f.year = year
}

func (f *Book) setSize(size int32) {
	f.size = size
}

func (f *Book) setRate(rate float32) {
	f.rate = rate
}

// Реализуйте методы для установки и получения полей структуры.

func (f Book) getID() int32 {
	return f.id
}

func (f Book) getTitle() string {
	return f.title
}

func (f Book) getAuthor() string {
	return f.author
}

func (f Book) getYear() int32 {
	return f.year
}

func (f Book) getSize() int32 {
	return f.size
}

func (f Book) getRate() float32 {
	return f.rate
}

// Реализуйте структуру с методом позволяющим сравнивать книги по полям Year, Size, Rate.
// Выбор режима сравнения задается в конструкторе структуры через перечисление (enum).
// Метод принимает 2 книги и выдает true если первый аргумент больше второго и false если наоборот.

func fieldsConstructor(byYear, bySize, byRate int) *Fields {
	return &Fields{byYear: byYear, bySize: bySize, byRate: byRate}
}

func (p *Fields) isYearEqual() {
	if p.byYear != 0 {
		fmt.Println("We will compare by year")
	}
}

func main() {
	// Place your code here.

	firstBook := Book{}
	secondBook := Book{}
	firstBook.setID(1)
	firstBook.setTitle("Harry Potter and the Philosopher's Stone")
	firstBook.setAuthor("J. K. Rowling")
	firstBook.setYear(1997)
	firstBook.setSize(300)
	firstBook.setRate(0.7)
	fmt.Println(firstBook)
	secondBook.setID(2)
	secondBook.setTitle("Harry Potter and the Chamber of Secrets")
	secondBook.setAuthor("J. K. Rowling")
	secondBook.setYear(1998)
	secondBook.setSize(300)
	secondBook.setRate(0.5)
	fmt.Println(secondBook)

	fmt.Println(firstBook.getID())
	fmt.Println(firstBook.getTitle())
	fmt.Println(firstBook.getAuthor())
	fmt.Println(firstBook.getYear())
	fmt.Println(firstBook.getSize())
	fmt.Println(firstBook.getRate())

	field := fieldsConstructor(1, 0, 0)
	field.isYearEqual()
}
