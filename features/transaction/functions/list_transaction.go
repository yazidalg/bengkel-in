package transaction

import (
	tStruct "bengkelin/features/transaction/structs"
	"fmt"
)

func ListTransaction(transactions tStruct.ArrTransaction) {
	var input string
	
	ShowListTransactionMenu()

	// Meminta input dari user untuk memilih menu
	fmt.Print("→ Masukan kode menu : ")
	fmt.Scan(&input)

	for input != "0" {
		
		// Mengecek ketersediaan menu
		if input == "1" {
			ListTransaction(transactions)
		} else if input == "2" {
			fmt.Println("Anda berada di menu 2")
		} else if input == "3" {
			fmt.Println("Anda berada di menu 3")
		} else {
			fmt.Println("Yah menu ga tersedia nih 😩")
		}

		// Meminta input dari user untuk memilih menu
		fmt.Print("→ Masukan kode menu : ")
		fmt.Scan(&input)
	}
}

func listAllTransaction(transactions tStruct.ArrTransaction) {
	for i := 0; i < transactions.N; i++ {
		fmt.Print(i + 1, ". ")
		fmt.Println("ID:", transactions.Data[i].Id)
		fmt.Printf("- Tanggal: %s-%s-%s", transactions.Data[i].Year, transactions.Data[i].Month, transactions.Data[i].Date)
		// fmt.Println("- Pembeli:", transactions.Data[i].Customer.Name,
		// 	"- Telp:", transactions.Data[i].Phone,
		// 	"- Alamat:", transactions.Data[i].Address)
	}
}
