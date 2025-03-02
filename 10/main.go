package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float32
}

type ShapeWithAria interface {
	Area() float32
}

type ShapeWithPerimeter interface {
	Perimeter() float32
}

type Square struct {
	ShapeWithAria
	ShapeWithPerimeter
}

func (s Square) Area() float32 {
	return s.sideLength * s.sideLength
}

type Circle struct {
	radius float32
}

func (c Circle) Perimeter() float32 {
	return c.radius * 4
}

func (c Circle) Area() float32 {
	return c.radius * c.radius * math.Pi
}

type Rectangle struct{}

func main() {
	square := Square{5}
	circle := Circle{8}

	printShapeAria(square)
	printShapeAria(circle)

	//printInterface(square)
	//printInterface(circle)
	printInterface("lol")
}

func printShapeAria(shape Shape) {
	fmt.Println("ShapeAria ===>", shape.Area())
}

func printInterface(i interface{}) {
	switch t := i.(type) {
	case int:
		fmt.Println("int =", t)
	case bool:
		fmt.Println("bool =", t)
	default:
		fmt.Println("unknown type", t)
	}

	str, ok := i.(string)
	if !ok {
		fmt.Println("string =", str)
	}
	fmt.Println("str", str)
	fmt.Printf("%+v/n", i)
}
