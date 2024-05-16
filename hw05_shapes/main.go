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
	circleFirst := circle{radius: 5}
	rectangleFirst := rectangle{width: 10, height: 5}
	triangleFirst := triangle{base: 8, height: 6}

	areaCircle, err := calculateArea(circleFirst)
	if err != nil {
		fmt.Println(err)
	}
	areaRectangle, err := calculateArea(rectangleFirst)
	if err != nil {
		fmt.Println(err)
	}
	areaTriangle, err := calculateArea(triangleFirst)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Круг: радиус %.2f. Площадь: %.2f\n", circleFirst.radius, areaCircle)
	fmt.Printf("Прямоугольник: ширина %.2f, высота %.2f. Площадь: %.2f\n",
		rectangleFirst.width, rectangleFirst.height, areaRectangle)
	fmt.Printf("Треугольник: основание %.2f, высота %.2f. Площадь: %.2f\n",
		triangleFirst.base, triangleFirst.height, areaTriangle)
}
