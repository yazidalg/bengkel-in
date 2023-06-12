package common

import "fmt"

/**
	Dibuat oleh Indra Mahesa 1302220067
**/
func ShowCustomConfirmationMessage(text string) bool {
	var inputRune rune
	var inputString string

	fmt.Println(text)

	for string(inputRune) != "\n" {
		fmt.Scanf("%c", &inputRune)

		if string(inputRune) != "\n" {
			inputString = inputString + string(inputRune)
		}
	}
	
	if inputString == "batal" {
		return false
	}
	
	return true
}