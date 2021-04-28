package main

import "fmt"

type User struct {
	name string
	age  uint8
}

func (u User) save() {
	fmt.Printf("O usuário '%s' de %d anos, foi salvo com sucesso!", u.name, u.age)
}

// Working with pointers in methods
func (u *User) incrementAge() {
	u.age++
}

func main() {
	user := User{"Leonardo", 22}

	user.incrementAge()

	user.save()
}
