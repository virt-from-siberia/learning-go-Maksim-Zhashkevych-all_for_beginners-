package main

import (
	"fmt"
	"reflect"
)

func main() {

	// messgae := "Hello world"

	var secondMessage int

	fmt.Println(reflect.TypeOf(secondMessage))

	var message string
	var number float64
	var b bool

	newMesage := []byte("asd")

	var a rune = 'a'

	fmt.Println(message)
	fmt.Println(number)
	fmt.Println(b)
	fmt.Println(newMesage)
	fmt.Println(a)

	a2, b2, c2 := 1, 2, 3

	a2, b2 = b2, a2

	fmt.Println(a2)
	fmt.Println(b2)
	fmt.Println(c2)
}
