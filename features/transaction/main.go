package transaction

import (
	"bengkelin/common"
	cStruct "bengkelin/features/customer/structs"
	sStruct "bengkelin/features/sparepart/structs"
	tFunc "bengkelin/features/transaction/functions"
	tStruct "bengkelin/features/transaction/structs"
	"fmt"
)

/**
	Dibuat oleh Indra Mahesa 1302220067
**/
func Main(customers cStruct.ArrCustomer, transactions *tStruct.ArrTransaction, spareparts *sStruct.ArrSparepart) {
	common.ResetConsole()

	var input string

	// Menampilkan main menu
	tFunc.ShowTransactionMenu()

	// Meminta input dari user untuk memilih menu
	fmt.Print("→ Masukan kode menu : ")
	fmt.Scan(&input)

	for input != "0" {
		// Mengecek ketersediaan menu
		if input == "1" {
			tFunc.ListTransaction(*transactions, *spareparts)
			tFunc.ShowTransactionMenu()
		} else if input == "2" {
			tFunc.CreateTransaction(customers, transactions, spareparts)
			tFunc.ShowTransactionMenu()
		} else if input == "3" {
			tFunc.EditTransaction(transactions)
			tFunc.ShowTransactionMenu()
		} else if input == "4" {
			tFunc.DeleteTransaction(transactions)
			tFunc.ShowTransactionMenu()
		} else {
			fmt.Println("Yah menu ga tersedia nih 😩")
		}

		// Meminta input dari user untuk memilih menu
		fmt.Print("→ Masukan kode menu : ")
		fmt.Scan(&input)
	}

	common.ResetConsole()
}
