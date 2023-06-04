package common

import "fmt"

func ShowConfirmationMessage() bool {
	var inputRune rune
	var inputString string

	fmt.Println("Tekan enter untuk konfirmasi atau ketik 'batal' untuk membatalkan...")

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