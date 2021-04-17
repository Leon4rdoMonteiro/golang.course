package main

import (
	"fmt"

	"github.com/badoux/checkmail"
)

func main() {
	// auxiliar.Write("Escrevendo algo na tela...")

	error := checkmail.ValidateFormat("leonardo@email.com")

	fmt.Println((error))
}
