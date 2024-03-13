package main

import (
	"errors"
	"fmt"
)

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

type BadShape struct{}

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

func calculateArea(s any) (float64, error) {
	shape, isOK := s.(Shape)
	if !isOK {
		return 0, errors.New("переданный объект не реализует интерфейс Shape")
	}
	return shape.Area()
}

func main() {
	// Place your code here.
	circle := Circle{Radius: -2}
	fmt.Println("круг: радиус", circle.Radius)

	circleArea, err := calculateArea(&circle)
	if err != nil {
		fmt.Println("не могу посчитать площадь: ", err)
	} else {
		fmt.Println("площадь: ", circleArea)
	}

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

	badshape := BadShape{}
	BadShapeArea, err := calculateArea(&badshape)
	if err != nil {
		fmt.Println("не могу посчитать площадь: ", err)
	} else {
		fmt.Println("площадь: ", BadShapeArea)
	}
}
