package common

import "fmt"

func InputMultipleString(value *string) {
	var inputRune rune
	
	*value = ""

	for string(inputRune) != "." {
		fmt.Scanf("%c", &inputRune)

		if string(inputRune) != "." {
			if string(inputRune) == "\n" {
				*value = *value + " "
			} else {
				*value = *value + string(inputRune)
			}
		}
	}
}