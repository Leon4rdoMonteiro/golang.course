package auxiliar

import "fmt"

// Write function, writes a text on the screen
func Write(text string) {
	fmt.Println(text)
	Write2("Escrevendo na tela 2...")
}
