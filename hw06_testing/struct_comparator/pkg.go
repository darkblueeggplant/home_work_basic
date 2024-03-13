package struct_comparator

type Book struct {
	Id     int32
	Title  string
	Author string
	Year   int32
	Size   int32
	Rate   float32
}

type Comparator struct {
	fieldCompare string
}

// Реализуйте методы для установки и получения полей структуры.

func (f *Book) SetID(id int32) {
	f.Id = id
}

func (f *Book) SetTitle(title string) {
	f.Title = title
}

func (f *Book) SetAuthor(author string) {
	f.Author = author
}

func (f *Book) SetYear(year int32) {
	f.Year = year
}

func (f *Book) SetSize(size int32) {
	f.Size = size
}

func (f *Book) SetRate(rate float32) {
	f.Rate = rate
}

// Реализуйте методы для установки и получения полей структуры.

// func (f Book) ID() int32 {
// 	return f.id
// }

// func (f Book) Title() string {
// 	return f.title
// }

// func (f Book) Author() string {
// 	return f.author
// }

// func (f Book) Year() int32 {
// 	return f.year
// }

// func (f Book) Size() int32 {
// 	return f.size
// }

// func (f Book) Rate() float32 {
// 	return f.rate
// }

// func isYearEqual(x int32, y int32) bool {
// 	if x > y {
// 		return x > y
// 	}
// 	return false
// }

// func isSizeEqual(x int32, y int32) bool {
// 	if x > y {
// 		return x > y
// 	}
// 	return false
// }

// func isRateEqual(x float32, y float32) bool {
// 	if x > y {
// 		return x > y
// 	}
// 	return false
// }

// Реализуйте структуру с методом позволяющим сравнивать книги по полям Year, Size, Rate.
// Выбор режима сравнения задается в конструкторе структуры через перечисление (enum).
// Метод принимает 2 книги и выдает true если первый аргумент больше второго и false если наоборот.

// func NewComparator(fieldCompare string) *Comparator {
// 	return &Comparator{fieldCompare: fieldCompare}
// }

// func (p *Comparator) isEqual(x, y Book) bool {
// 	switch p.fieldCompare {
// 	case "year":
// 		return isYearEqual(x.year, y.year)

// 	case "size":
// 		return isSizeEqual(x.size, y.size)

// 	case "rate":
// 		return isRateEqual(x.rate, y.rate)

// 	default:
// 		return false
// 	}
// }

// func main() {
// 	// Place your code here.

// firstBook := Book{}
// secondBook := Book{}
// 	firstBook.SetID(1)
// 	firstBook.SetTitle("Harry Potter and the Philosopher's Stone")
// 	firstBook.SetAuthor("J. K. Rowling")
// 	firstBook.SetYear(1997)
// 	firstBook.SetSize(300)
// 	firstBook.SetRate(0.7)
// 	fmt.Println(firstBook)
// 	secondBook.SetID(2)
// 	secondBook.SetTitle("Harry Potter and the Chamber of Secrets")
// 	secondBook.SetAuthor("J. K. Rowling")
// 	secondBook.SetYear(1998)
// 	secondBook.SetSize(300)
// 	secondBook.SetRate(0.5)
// 	fmt.Println(secondBook)

// 	fmt.Println(firstBook.ID())
// 	fmt.Println(firstBook.Title())
// 	fmt.Println(firstBook.Author())
// 	fmt.Println(firstBook.Year())
// 	fmt.Println(firstBook.Size())
// 	fmt.Println(firstBook.Rate())

// 	comparator := NewComparator("year")
// 	fmt.Println(comparator.isEqual(firstBook, secondBook))

// 	comparator = NewComparator("rate")
// 	fmt.Println(comparator.isEqual(firstBook, secondBook))
// }
