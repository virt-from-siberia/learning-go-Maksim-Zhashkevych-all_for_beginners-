package main

import "fmt"

type Age int

func (a Age) isAdult() bool {
	return a > 18
}

type User struct {
	name   string
	age    Age
	sex    string
	weight int
	height int
}

func (u *User) printUserInfo(name string) {
	u.name = name
	fmt.Println(u.name, u.age, u.weight, u.height)
}

func (u User) getName() string {
	return u.name
}

func (u *User) setName(name string) {
	u.name = name
}

type DumbDatabase struct {
	m map[string]string
}

func newBD() *DumbDatabase {
	return &DumbDatabase{make(map[string]string)}
}

func newUser(name, sex string, age, weight, height int) User {
	return User{
		name:   name,
		age:    Age(age),
		sex:    sex,
		weight: weight,
		height: height,
	}
}

func main() {
	user1 := newUser("Vasua", "male", 75, 22, 27)
	//user2 := User{"Alex", 22, "male", 25, 1}

	//user1.sex = "women"

	//fmt.Printf("%+v/n", user1)
	//fmt.Printf("%+v/n", user2)
	//
	//user1.printUserInfo("Kolya")
	//user2.printUserInfo("Maxim")

	user1.setName("new Alex")

	fmt.Println(user1.age.isAdult())
	fmt.Println(user1.getName())
	//fmt.Println(user1.getName())
}
