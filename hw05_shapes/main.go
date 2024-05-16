package main

import (
	"errors"
	"fmt"
	"math"
)

type Shape interface {
	area() (float64, error)
}

type circle struct {
	radius float64
}

type rectangle struct {
	width, height float64
}

type triangle struct {
	base, height float64
}

func (c circle) area() (float64, error) {
	if c.radius < 0 {
		return 0, errors.New("радиус не может быть отрицательным")
	}
	return math.Pi * c.radius * c.radius, nil
}

func (r rectangle) area() (float64, error) {
	if r.width < 0 || r.height < 0 {
		return 0, errors.New("ширина и высота не могут быть отрицательными")
	}
	return r.width * r.height, nil
}

func (t triangle) area() (float64, error) {
	if t.base < 0 || t.height < 0 {
		return 0, errors.New("основание и высота не могут быть отрицательными")
	}
	return t.base * t.height / 2, nil
}

func calculateArea(s Shape) (any, error) {
	return s.area()
}

func main() {
	circle := circle{radius: 5}
	rectangle := rectangle{width: 10, height: 5}
	triangle := triangle{base: 8, height: 6}

	areaCircle, err := calculateArea(circle)
	if err != nil {
		fmt.Println(err)
	}
	areaRectangle, err := calculateArea(rectangle)
	if err != nil {
		fmt.Println(err)
	}
	areaTriangle, err := calculateArea(triangle)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Круг: радиус %.2f. Площадь: %.2f\n", circle.radius, areaCircle)
	fmt.Printf("Прямоугольник: ширина %.2f, высота %.2f. Площадь: %.2f\n",
		rectangle.width, rectangle.height, areaRectangle)
	fmt.Printf("Треугольник: основание %.2f, высота %.2f. Площадь: %.2f\n", triangle.base, triangle.height, areaTriangle)
}
