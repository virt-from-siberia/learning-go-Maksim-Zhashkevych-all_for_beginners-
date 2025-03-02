package main

import (
	"errors"
	"fmt"
)

func main() {

	messages := make([]string, 100)
	fmt.Println(len(messages))
	fmt.Println(cap(messages))

	messages = append(messages, "101")

	fmt.Println(len(messages))
	fmt.Println(cap(messages))

	fmt.Println(messages)

}

func printMessages(messages []string) error {
	if len(messages) == 0 {
		return errors.New("empty array")
	}
	fmt.Println(messages)
	return nil
}
