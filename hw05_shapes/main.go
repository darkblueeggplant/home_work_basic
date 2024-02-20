package main

import "fmt"

type Shape interface {
	calculateArea() (float64, error)
}

type Circle struct {
	radius float64
}
type Rectangle struct {
	width  float64
	height float64
}
type Triangle struct {
	height float64
	side   float64
}

type BadShape struct{}

func (b BadShape) calculateArea() (float64, error) {
	return 0, fmt.Errorf("переданный объект не является фигурой")
}

func (c Circle) calculateArea() (float64, error) {
	if c.radius < 0 {
		return 0, fmt.Errorf("задан отрицательный радиус")
	}
	return 3.14 * c.radius * c.radius, nil
}

func (r Rectangle) calculateArea() (float64, error) {
	if r.height < 0 || r.width < 0 {
		return 0, fmt.Errorf("задано отрицательное значение одной или обеих сторон")
	}
	return r.width * r.height, nil
}

func (t Triangle) calculateArea() (float64, error) {
	if t.height < 0 || t.side < 0 {
		return 0, fmt.Errorf("задано отрицательное значение основания или высоты")
	}
	return 0.5 * t.height * t.side, nil
}

func calculateArea(shape Shape) (float64, error) {
	// return shape.calculateArea()
	area, err := shape.calculateArea()
	if err != nil {
		return 0, err
	}
	return area, nil
}

func main() {
	// Place your code here.
	circle := Circle{radius: 2}
	fmt.Println("круг: радиус", circle.radius)

	circleArea, err := calculateArea(circle)
	if err != nil {
		fmt.Println("не могу посчитать площадь: ", err)
	} else {
		fmt.Println("площадь: ", circleArea)
	}

	circleBad := Circle{radius: -3}
	fmt.Println("круг: радиус", circleBad.radius)

	circleBadArea, err := calculateArea(circleBad)
	if err != nil {
		fmt.Println("не могу посчитать площадь: ", err)
	} else {
		fmt.Println("площадь: ", circleBadArea)
	}

	rectangle := Rectangle{width: 3, height: 2}
	fmt.Println("прямоугольник: ширина", rectangle.width, "высота:", rectangle.height)

	rectangleArea, err := calculateArea(rectangle)
	if err != nil {
		fmt.Println("не могу посчитать площадь: ", err)
	} else {
		fmt.Println("площадь: ", rectangleArea)
	}

	triangle := Triangle{height: 10, side: 3}
	fmt.Println("треуголльник: высота", triangle.height, "основание:", triangle.side)

	triangleArea, err := calculateArea(triangle)
	if err != nil {
		fmt.Println("не могу посчитать площадь: ", err)
	} else {
		fmt.Println("площадь: ", triangleArea)
	}

	invalidShape := BadShape{}
	invalidShapeArea, err := calculateArea(invalidShape)
	if err != nil {
		fmt.Println("не могу посчитать площадь: ", err)
	} else {
		fmt.Println("площадь: ", invalidShapeArea)
	}
}
