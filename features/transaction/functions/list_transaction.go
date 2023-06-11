package transaction

import (
	"bengkelin/common"
	tStruct "bengkelin/features/transaction/structs"
	"fmt"
)

func ListTransaction(transactions tStruct.ArrTransaction) {
	common.ResetConsole()

	var input string
	
	ShowListTransactionMenu()

	// Meminta input dari user untuk memilih menu
	fmt.Print("→ Masukan kode menu : ")
	fmt.Scan(&input)

	for input != "0" {
		
		// Mengecek ketersediaan menu
		if input == "1" {
			common.ResetConsole()
			fmt.Println("=======================================================================================")
			fmt.Println("                                    List Transaksi                                     ")
			fmt.Println("=======================================================================================")
			fmt.Println()
			listAllTransaction(transactions)
			ShowListTransactionMenu()
		} else if input == "2" {
			fmt.Println("Anda berada di menu 2")
		} else if input == "3" {
			common.ResetConsole()
			fmt.Println("=======================================================================================")
			fmt.Println("                       Lihat Pelanggan Berdasarkan Pembelian Periode                   ")
			fmt.Println("=======================================================================================")
			fmt.Println()
			listCustomerBySparepart(transactions)
		} else {
			fmt.Println("Yah menu ga tersedia nih 😩")
		}

		// Meminta input dari user untuk memilih menu
		fmt.Print("→ Masukan kode menu : ")
		fmt.Scan(&input)
	}

	common.ResetConsole()
}

func listCustomerBySparepart(transactions tStruct.ArrTransaction) {
	
}

func listAllTransaction(transactions tStruct.ArrTransaction) {
	for i := 0; i < transactions.N; i++ {
		fmt.Print(i + 1, ". ")
		fmt.Print("ID: ", transactions.Data[i].Id)
		fmt.Println()
		fmt.Printf("   Tanggal: %v-%v-%v ", transactions.Data[i].Year, transactions.Data[i].Month, transactions.Data[i].Date)
		fmt.Println()
		fmt.Printf("   Customer: %v ", transactions.Data[i].Customer.Name)
		fmt.Println()
		fmt.Printf("   Tarif: Rp. %v ", transactions.Data[i].Price)
		fmt.Println()
		fmt.Println()
		fmt.Println("- - - - - - - - - - - - - - - - ")
		fmt.Println()
	}
	common.ShowEndAction(1)
	common.ResetConsole()
}
