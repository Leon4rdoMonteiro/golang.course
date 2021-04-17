package main

import "fmt"

type User struct {
	name    string
	email   string
	age     uint8
	address Address
}

type Address struct {
	street string
	number uint8
}

func main() {
	fmt.Println("Structs file")

	var user User

	user.name = "Leonardo"
	user.email = "leonardo.monteiro@gmail.com"
	user.age = 23

	fmt.Println(user)

	// Using typing inference
	userAddress := Address{"Rua do Bar da Deuza", 186}
	user2 := User{"Leonardo", "leonardo.monteiro@gmail.com", 23, userAddress}

	fmt.Println(user2)

	// Fill only specific values

	user3 := User{name: "Leonardo", age: 23}

	fmt.Println(user3)
}
