package main

import "fmt"

var msg string

func init() {
	msg = "omg init"
}

func main() {
	fmt.Println(msg)
	// func() {
	// 	fmt.Println("Анонимная функция")
	// }()

	inc := increment()

	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
}

func increment() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}
