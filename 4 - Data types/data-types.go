package main

import (
	"errors"
	"fmt"
)

func main() {
	// INTEGERS
	const int8Bits int8 = 127
	const int16Bits int16 = 32767
	const int32Bits int32 = 2147483647
	const int64Bits int64 = 9223372036854775807
	const intByArchitecture int = 9223372036854775807 // Defined by machine architecture 32 or 64 bits / used in inference typing
	const unsignedInt uint16 = 65535                  // Not accept negative numbers

	const runeType rune = 2147483645 // Same of int32
	const byteType byte = 255        // Same of uint8

	// FLOATS

	// "float" type not exists as a "int" type
	const float32Bits float32 = 12345.2

	fmt.Println(int8Bits)

	// STRINGS

	// Use always "" to strings
	var stringType string = "Text"
	fmt.Println(stringType)

	// Chart type not exists on GOlang
	char := 'A'

	// Will print the ASC table number for the character
	fmt.Println(char)

	// BOOLEAN TYPE
	var boolean bool = true
	fmt.Println(boolean)

	// ERROR Type
	var errorType error = errors.New("Internal error")
	fmt.Println(errorType)

	// Default initial value for each type
	var zeroValue int
	fmt.Println(zeroValue)

	var emptyString string
	fmt.Println(emptyString)

	var boolFalse bool
	fmt.Println(boolFalse)

	var errorNil error
	fmt.Println(errorNil)
}
