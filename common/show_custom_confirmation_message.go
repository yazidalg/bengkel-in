package common

import "fmt"

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