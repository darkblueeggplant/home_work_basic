package main

import "fmt"

type Shape interface {
	Area() (float64, error)
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Triangle struct {
	Height float64
	Side   float64
}

// type BadShape struct{}

// func (b BadShape) calculateArea() (float64, error) {
// 	return 0, fmt.Errorf("переданный объект не является фигурой")
// }

func (c *Circle) Area() (float64, error) {
	if c.Radius < 0 {
		return 0, fmt.Errorf("задан отрицательный радиус")
	}
	return 3.14 * c.Radius * c.Radius, nil
}

func (r *Rectangle) Area() (float64, error) {
	if r.Height < 0 || r.Width < 0 {
		return 0, fmt.Errorf("задано отрицательное значение одной или обеих сторон")
	}
	return r.Width * r.Height, nil
}

func (t *Triangle) Area() (float64, error) {
	if t.Height < 0 || t.Side < 0 {
		return 0, fmt.Errorf("задано отрицательное значение основания или высоты")
	}
	return 0.5 * t.Height * t.Side, nil
}

func calculateArea(s Shape) (float64, error) {
	area, err := s.Area()
	if err != nil {
		return 0, err
	}
	return area, nil
}

func main() {
	// Place your code here.
	circle := Circle{Radius: 2}
	fmt.Println("круг: радиус", circle.Radius)

	circleArea, err := calculateArea(&circle)
	if err != nil {
		fmt.Println("не могу посчитать площадь: ", err)
	} else {
		fmt.Println("площадь: ", circleArea)
	}

	// circleBad := Circle{Radius: -3}
	// fmt.Println("круг: радиус", circleBad.Radius)

	// circleBadArea, err := calculateArea(circleBad)
	// if err != nil {
	// 	fmt.Println("не могу посчитать площадь: ", err)
	// } else {
	// 	fmt.Println("площадь: ", circleBadArea)
	// }

	rectangle := Rectangle{Width: 3, Height: 2}
	fmt.Println("прямоугольник: ширина", rectangle.Width, "высота:", rectangle.Height)

	rectangleArea, err := calculateArea(&rectangle)
	if err != nil {
		fmt.Println("не могу посчитать площадь: ", err)
	} else {
		fmt.Println("площадь: ", rectangleArea)
	}

	triangle := Triangle{Height: 10, Side: 3}
	fmt.Println("треуголльник: высота", triangle.Height, "основание:", triangle.Side)

	triangleArea, err := calculateArea(&triangle)
	if err != nil {
		fmt.Println("не могу посчитать площадь: ", err)
	} else {
		fmt.Println("площадь: ", triangleArea)
	}

	// invalidShape := BadShape{}
	// invalidShapeArea, err := calculateArea(invalidShape)
	// if err != nil {
	// 	fmt.Println("не могу посчитать площадь: ", err)
	// } else {
	// 	fmt.Println("площадь: ", invalidShapeArea)
	// }
}
