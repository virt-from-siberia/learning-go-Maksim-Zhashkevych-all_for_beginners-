package main

import "fmt"

func main() {
	users := map[string]int{
		"Robert": 34,
		"Alex":   15,
	}

	fmt.Println(users)
	fmt.Println(users["Robert"])
	fmt.Println(users["Lol"])

	age, exists := users["Robert"]
	if exists {
		fmt.Println("Exists")
	} else {
		fmt.Println("Not exists")
	}

	fmt.Println(age)

	for k, v := range users {
		fmt.Println(k, v)
	}
	users["Sergey"] = 21

	for k, v := range users {
		fmt.Println(k, v)
	}

	delete(users, "Robert")

	for k, v := range users {
		fmt.Println(k, v)
	}

	var customers map[string]int
	customers = make(map[string]int)

	customers["Vasua"] = 19

	fmt.Println(customers["Vasua"])
}
