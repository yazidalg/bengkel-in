package transaction

import (
	cStruct "bengkelin/features/customer/structs"
	sStruct "bengkelin/features/sparepart/structs"
	tFunc "bengkelin/features/transaction/functions"
	tStruct "bengkelin/features/transaction/structs"
	"fmt"
)

func Main(customers cStruct.ArrCustomer, transactions *tStruct.ArrTransaction, spareparts *sStruct.ArrSparepart) {
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
			tFunc.CreateTransaction(customers, transactions, spareparts)
			tFunc.ShowTransactionMenu()
		} else if input == "5" {
			fmt.Println("Yah menu ga tersedia nih 😩")
		} else if input == "6" {
			tFunc.DeleteTransaction(transactions)
		}

		// Meminta input dari user untuk memilih menu
		fmt.Print("→ Masukan kode menu : ")
		fmt.Scan(&input)
	}
}
