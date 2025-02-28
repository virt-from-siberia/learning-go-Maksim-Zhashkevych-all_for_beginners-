package main

import "fmt"

func main() {
	var p *int
	fmt.Println(p)

	message := "Скоро я стану ninja"
	changeMessage(&message)
	fmt.Println(message)
}

func changeMessage(message *string) {
	*message += " (из функции printMessage())"
	// fmt.Println(message)

}
