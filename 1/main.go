package main

import (
	"errors"
	"fmt"
)

func main() {

	// print("Вызо 1")

	// print("Вызо 2")
	// print("Вызо 3")

	// message, entered fmt.Println(sayHello("Alex", 39))

	message, entered, err := enterTheClub(15)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(enterTheClub(14))
	fmt.Println(message)
	fmt.Println(entered)
	fmt.Println(err)

	predictionResult, predictionErr := prediction("lol")

	fmt.Println(predictionResult)
	fmt.Println(predictionErr)

}

// func print(message string) {
// 	fmt.Println(message)
// }

func sayHello(name string, age int) string {
	return fmt.Sprintf("Привет %s Тебе %d лет", name, age)
}

func enterTheClub(age int) (string, bool, error) {
	if age >= 18 && age <= 45 {
		response := "Входи"
		return response, true, nil
	}
	return "Вход запрещен", false, errors.New("you are too young")
}

func prediction(dayOfWeek string) (string, error) {

	switch dayOfWeek {
	case "Пн":
		return "Хорошего начала недели!", nil
	case "Вт":
		return "Хорошего начала недели!", nil
	default:
		return "Неверный день недели", errors.New("invalid day of the week")
	}

}
