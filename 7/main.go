package main

import "fmt"

func main() {

	defer handlePanic()

	messages := []string{
		"Hello World",
		"Hello World",
		"Hello World",
		"Hello World",
	}

	messages[4] = "error"

	fmt.Println("main")
}

func handlePanic() {
	//fmt.Println("printMessage()")
	if r := recover(); r != nil {
		fmt.Println("recover", r)
	}

}
