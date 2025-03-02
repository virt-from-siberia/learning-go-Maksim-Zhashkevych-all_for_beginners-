package main

import "fmt"

func main() {
	message := []string{
		"Message 1",
		"Message 2",
		"Message 3",
	}

	for i := range message {
		fmt.Println(message[i])
	}

	for _, value := range message {
		fmt.Println(value)
	}

	counter := 0
	for counter < 10 {
		counter++
		fmt.Println(counter)
	}

}
