package main

import "fmt"

type Shape interface {
	calculateArea() float64
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

func (c Circle) calculateArea() float64 {
	return 3.14 * c.radius * c.radius
}

func (r Rectangle) calculateArea() float64 {
	return r.width * r.height
}

func (t Triangle) calculateArea() float64 {
	return 0.5 * t.height * t.side
}

func calculateArea(shape Shape) float64 {
	return shape.calculateArea()
}

func main() {
	// Place your code here.
	circle := Circle{radius: 2}
	fmt.Println("круг: радиус", circle.radius)
	fmt.Println("площадь круга:", calculateArea(circle))

	rectangle := Rectangle{width: 3, height: 2}
	fmt.Println("прямоугольник: ширина", rectangle.width, "высота:", rectangle.height)
	fmt.Println("площадь прямоугольника:", calculateArea(rectangle))

	triangle := Triangle{height: 10, side: 3}
	fmt.Println("треуголльник: высота", triangle.height, "основание:", triangle.side)
	fmt.Println("площадь треугольника:", calculateArea(triangle))

}
