package main

import (
	"testing"

	"github.com/darkblueeggplant/home_work_basic/hw06_testing/chessboard"
	"github.com/darkblueeggplant/home_work_basic/hw06_testing/struct_comparator"
)

// func TestPrintSomething(t *testing.T) {
// 	const str, want = "Hello", "Hello"
// 	got := chessboard.PrintSomething(str)
// 	if got != want {
// 		t.Errorf("something wrong: got %v, want %v", got, want)
// 	}
// }

// func Test(t *testing.T) {
// 	testCases := []struct {
// 		desc      string
// 		str, want string
// 	}{
// 		{
// 			desc: "chessboard.PrintSomething",
// 			str:  "Hello",
// 			want: "Hello",
// 		},
// 	}
// 	for _, tC := range testCases {
// 		t.Run(tC.desc, func(t *testing.T) {
// 			got := chessboard.PrintSomething(tC.str)
// 			if got != tC.want {
// 				t.Errorf("something wrong: got %v, want %v", got, tC.want)
// 			}
// 		})
// 	}
// }

func Test(t *testing.T) {
	testCasesC := []struct {
		desc string
		size int
		want string
	}{
		{
			desc: "chessboard.PrintDesk",
			size: 2,
			want: "# \n #\n",
		},
	}

	for _, tC := range testCasesC {
		t.Run(tC.desc, func(t *testing.T) {
			got := chessboard.PrintDesk(tC.size)
			if got != tC.want {
				t.Errorf("something wrong: got %v, want %v", got, tC.want)
			}

		})
	}
	testCasesSCS := []struct {
		desc   string
		id     int32
		title  string
		author string
		year   int32
		size   int32
		rate   float32
		want   struct_comparator.Book
	}{
		{
			desc:   "struct_comparator.Setters",
			id:     1,
			title:  "Harry Potter and the Philosopher's Stone",
			author: "J. K. Rowling",
			year:   1997,
			size:   300,
			rate:   0.7,

			want: struct_comparator.Book{
				Id:     1,
				Title:  "Harry Potter and the Philosopher's Stone",
				Author: "J. K. Rowling",
				Year:   1997,
				Size:   300,
				Rate:   0.7,
			},
		},
	}

	for _, tC := range testCasesSCS {
		t.Run(tC.desc, func(t *testing.T) {
			testBookSet := struct_comparator.Book{}
			testBookSet.SetID(tC.id)
			testBookSet.SetTitle(tC.title)
			testBookSet.SetAuthor(tC.author)
			testBookSet.SetYear(tC.year)
			testBookSet.SetSize(tC.size)
			testBookSet.SetRate(tC.rate)
			Sgot := testBookSet
			if Sgot != tC.want {
				t.Errorf("something wrong: got %v, want %v", Sgot, tC.want)
			}

		})

	}
	//
	testCasesSCGID := []struct {
		desc string
		id   int32
		want int32
	}{
		{
			desc: "struct_comparator.GettersID",
			id:   1,
			want: 1,
		},
	}
	for _, tC := range testCasesSCGID {
		t.Run(tC.desc, func(t *testing.T) {
			testBookGet := struct_comparator.Book{
				Id:     tC.id,
				Title:  "",
				Author: "",
				Year:   0,
				Size:   0,
				Rate:   0,
			}
			SCGIdgot := testBookGet.ID()

			if SCGIdgot != tC.want {
				t.Errorf("something wrong: got %v, want %v", SCGIdgot, tC.want)
			}

		})
	}
	testCasesSCGTITLE := []struct {
		desc  string
		title string
		want  string
	}{
		{
			desc:  "struct_comparator.GettersTITLE",
			title: "Harry Potter and the Philosopher's Stone",
			want:  "Harry Potter and the Philosopher's Stone",
		},
	}
	for _, tC := range testCasesSCGTITLE {
		t.Run(tC.desc, func(t *testing.T) {
			testBookGet := struct_comparator.Book{
				Id:     0,
				Title:  tC.title,
				Author: "",
				Year:   0,
				Size:   0,
				Rate:   0,
			}
			SCGTitlegot := testBookGet.TITLE()

			if SCGTitlegot != tC.want {
				t.Errorf("something wrong: got %v, want %v", SCGTitlegot, tC.want)
			}

		})
	}
	testCasesSCIsEqual := []struct {
		desc string
		x    int32
		y    int32
		want bool
	}{
		{
			desc: "struct_comparator.IsYearEqual",
			x:    1987,
			y:    1986,
			want: true,
		},
		{
			desc: "struct_comparator.IsYearEqual",
			x:    1986,
			y:    1987,
			want: false,
		},
	}
	for _, tC := range testCasesSCIsEqual {
		t.Run(tC.desc, func(t *testing.T) {
			got := struct_comparator.IsYearEqual(tC.x, tC.y)
			if got != tC.want {
				t.Errorf("something wrong: got %v, want %v", got, tC.want)
			}

		})
	}
}

// type Book struct {
// 	id     int32
// 	title  string
// 	author string
// 	year   int32
// 	size   int32
// 	rate   float32
// }

// func TestStructComparator(t *testing.T) {
// 	testCases := []struct {
// 		desc string
// 		// Id	int32
// 		// Title string
// 		// Author string
// 		// Year int32
// 		// Size int32
// 		// Rate float32
// 		id   int32
// 		want struct_comparator.Book
// 	}{
// 		{
// 			desc: "struct_comparator.SetID",
// 			id:   2,
// 			want: struct_comparator.Book{
// 				Id:     1,
// 				Title:  "",
// 				Author: "",
// 				Year:   0,
// 				Size:   0,
// 				Rate:   0,
// 			},
// 		},
// 	}
// 	for _, tC := range testCases {
// 		t.Run(tC.desc, func(t *testing.T) {
// 			firstBook := struct_comparator.Book{}
// 			firstBook.SetID(tC.id)
// 			got := firstBook
// 			if got != tC.want {
// 				t.Errorf("something wrong: got %v, want %v", got, tC.want)
// 			}
// 		})
// 	}
// }

//
// func TestSetID(t *testing.T) {
// 	const id = 1
// 	firstBook := struct_comparator.Book{}
// 	firstBook.SetID(id)
// 	got := firstBook
// 	want := struct_comparator.Book{
// 		Id:     1,
// 		Title:  "",
// 		Author: "",
// 		Year:   0,
// 		Size:   0,
// 		Rate:   0,
// 	}

// 	if got != want {
// 		t.Errorf("alarm!")
// 	}

// }
//

// type Book struct {
// 	id     int32
// 	title  string
// 	author string
// 	year   int32
// 	size   int32
// 	rate   float32
// }

// func TestPrintDesk(t *testing.T) {
// 	const size = 2
// 	want := "# \n #\n"
// 	got := chessboard.PrintDesk(size)
// 	if got != want {
// 		t.Errorf("something wrong: got %v, want %v", got, want)
// 	}
// }

// func TestSetID(t *testing.T) {
// 	const id = 1
// 	firstBook := struct_comparator.Book{}
// 	firstBook.SetID(id)
// 	got := firstBook
// 	want
// got := firstBook.SetID(id)
// if got != want {
// 	t.Errorf("something wrong: got %v, want %v", got, want)
// }
// 	firstBook.SetID(1)
// firstBook.SetTitle("Harry Potter and the Philosopher's Stone")
// firstBook.SetAuthor("J. K. Rowling")
// firstBook.SetYear(1997)
// firstBook.SetSize(300)
// firstBook.SetRate(0.7)
// got := firstBook
// want := []struct {
// 	id     int32
// 	title  string
// 	author string
// 	year   int32
// 	size   int32
// 	rate   float32
// }{
// 	{
// 		id:     1,
// 		title:  "Harry Potter and the Philosopher's Stone",
// 		author: "J. K. Rowling",
// 		year:   1997,
// 		size:   300,
// 		rate:   0.7,
// 	},
// }
// want := struct_comparator.Book{1, "Harry Potter and the Philosopher's Stone", "J. K. Rowling", 1997, 300, 0.7}
// if got != want {
// 	t.Errorf("something wrong")
// }
// }
