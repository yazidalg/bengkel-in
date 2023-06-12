package common

import "fmt"

/**
	Dibuat oleh Indra Mahesa 1302220067
**/
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