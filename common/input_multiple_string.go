package common

import "fmt"

// Ini berfungsi untuk menginput string dengan spasi
func InputMultipleString(value *string) {
	var inputRune rune
	
	*value = ""

	for string(inputRune) != "\n" {
		fmt.Scanf("%c", &inputRune)

		if string(inputRune) != "\n" {
			*value = *value + string(inputRune)
		}
	}
}