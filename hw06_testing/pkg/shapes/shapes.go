package shapes

import (
	"errors"
	"math"
)

type Shape interface {
	Area() (float64, error)
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width, Height float64
}

type Triangle struct {
	Base, Height float64
}

func (c Circle) Area() (float64, error) {
	if c.Radius < 0 {
		return 0, errors.New("радиус не может быть отрицательным")
	}
	return math.Pi * c.Radius * c.Radius, nil
}

func (r Rectangle) Area() (float64, error) {
	if r.Width < 0 || r.Height < 0 {
		return 0, errors.New("ширина и высота не могут быть отрицательными")
	}
	return r.Width * r.Height, nil
}

func (t Triangle) Area() (float64, error) {
	if t.Base < 0 || t.Height < 0 {
		return 0, errors.New("основание и высота не могут быть отрицательными")
	}
	return t.Base * t.Height / 2, nil
}

func CalculateArea(s Shape) (float64, error) {
	return s.Area()
}
