package common

import "fmt"

/**
	Dibuat oleh Indra Mahesa 1302220067
**/
func ShowEndAction(repeat int) {
	var tempInput string

	fmt.Println()
	fmt.Println("Tekan enter untuk selesai...")


	for i := 0; i < repeat; i++ {
		fmt.Scanln(&tempInput)
	}
}