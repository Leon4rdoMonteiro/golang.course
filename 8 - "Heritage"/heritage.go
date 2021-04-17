package main

import "fmt"

type Person struct {
	first_name string
	last_name  string
	age        uint8
	height     float32
}

type Student struct {
	Person
	course     string
	university string
}

func main() {
	fmt.Println("'Heritage' - Golang")

	person := Person{"Leonardo", "Monteiro", 23, 1.64}
	fmt.Println(person)

	student := Student{person, "ADS", "UDF"}
	fmt.Println(student)
}
