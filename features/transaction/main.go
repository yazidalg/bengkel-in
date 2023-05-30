package transaction

import (
	tFunc "bengkelin/features/transaction/functions"
	tStruct "bengkelin/features/transaction/structs"
	"fmt"
)

func Main(transactions *tStruct.ArrTransaction) {
	var input string
	
	// Menampilkan main menu
	tFunc.ShowTransactionMenu()

	// Meminta input dari user untuk memilih menu
	fmt.Print("→ Masukan kode menu : ")
	fmt.Scan(&input)
	
	for input != "7" {
		
		// Mengecek ketersediaan menu
		if input == "1" {
			fmt.Println("Anda berada di menu 1")
		} else if input == "2" {
			fmt.Println("Anda berada di menu 2")
		} else if input == "3" {
			fmt.Println("Anda berada di menu 3")
		} else if input == "4" {
			if tFunc.CreateTransaction(transactions) {
				fmt.Println("Berhasil membuat transaksi")
			}
		} else {
			fmt.Println("Yah menu ga tersedia nih 😩")
		}

		// Meminta input dari user untuk memilih menu
		fmt.Print("→ Masukan kode menu : ")
		fmt.Scan(&input)
	}
}
