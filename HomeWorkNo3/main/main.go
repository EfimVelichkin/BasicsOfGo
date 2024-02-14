package main

import (
	"fmt"
	"math"
)

type Shape struct {
	Name string
}

func (s Shape) GetName() string {
	return s.Name
}

func (s Shape) Area() float64 {
	return 0.0
}

type Circle struct {
	Shape
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

type Rectangle struct {
	Shape
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	var EntRad float64
	fmt.Print("Введите радиус круга: ")
	fmt.Scan(&EntRad)

	var EntWid float64
	fmt.Print("Введите сторону A прямоугольника: ")
	fmt.Scan(&EntWid)

	var EntHeig float64
	fmt.Print("Введите сторону B прямоугольника: ")
	fmt.Scan(&EntHeig)

	circle := Circle{Shape{"Круг"}, EntRad}
	rectangle := Rectangle{Shape{"Прямоугольник"}, EntWid, EntHeig}

	fmt.Printf("%s: Площадь = %.2f\n", circle.GetName(), circle.Area())
	fmt.Printf("%s: Площадь = %.2f\n", rectangle.GetName(), rectangle.Area())
}
