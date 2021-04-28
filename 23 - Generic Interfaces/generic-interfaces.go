package main

import (
	"errors"
	"fmt"
)

func printAnything(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	printAnything(nil)
	printAnything(errors.New("Error"))
	printAnything("nil")
	printAnything(0)
	printAnything(0.0)
}
