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

type param struct {
	field string
}

// Реализуйте методы для установки и получения полей структуры.
func (f *Book) create(id int32, title string, author string, year int32, size int32, rate float32) {
	f.id = id
	f.title = title
	f.author = author
	f.year = year
	f.size = size
	f.rate = rate
}

// Реализуйте методы для установки и получения полей структуры.
func (f *Book) get() (int32, string, string, int32, int32, float32) {
	return f.id, f.title, f.author, f.year, f.size, f.rate
}

func setParam(field string) *param {
	return &param{field: field}
}

func isYearEqual(x int32, y int32) bool {
	if x > y {
		return x > y
	}
	return false
}

func isSizeEqual(x int32, y int32) bool {
	if x > y {
		return x > y
	}
	return false
}

func isRateEqual(x float32, y float32) bool {
	if x > y {
		return x > y
	}
	return false
}

// Реализуйте структуру с методом позволяющим сравнивать книги по полям Year, Size, Rate.
// Выбор режима сравнения задается в конструкторе структуры через перечисление (enum).
// Метод принимает 2 книги и выдает true если первый аргумент больше второго и false если наоборот.

func (p param) isEqual(x, y Book) bool {
	switch p.field {
	case "year":
		return isYearEqual(x.year, y.year)

	case "size":
		return isSizeEqual(x.size, y.size)

	case "rate":
		return isRateEqual(x.rate, y.rate)

	default:
		return false
	}
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

	// books := []Book{
	// 	{3, "Harry Potter and the Prisoner of Azkaban", "J. K. Rowling", 1999, 300, 0.25},
	// 	{4, "Harry Potter and the Goblet of Fire", "J. K. Rowling", 2000, 300, 1.0},
	// }
	// fmt.Println(books)

	fmt.Println(param.isEqual(*setParam("year"), firstBook, secondBook))
	fmt.Println(param.isEqual(*setParam("size"), firstBook, secondBook))
	fmt.Println(param.isEqual(*setParam("rate"), firstBook, secondBook))
	fmt.Println(param.isEqual(*setParam("blahblahblah"), firstBook, secondBook))
}
