package common

import "fmt"

/**
	Dibuat oleh Indra Mahesa 1302220067
**/
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