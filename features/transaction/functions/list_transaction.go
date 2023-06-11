package transaction

import (
	"bengkelin/common"
	cFunc "bengkelin/features/customer/functions"
	sFunc "bengkelin/features/sparepart/functions"
	sStruct "bengkelin/features/sparepart/structs"
	tStruct "bengkelin/features/transaction/structs"
	"fmt"
)

func ListTransaction(transactions tStruct.ArrTransaction, spareparts sStruct.ArrSparepart) {
	common.ResetConsole()

	var input string

	ShowListTransactionMenu()

	// Meminta input dari user untuk memilih menu
	fmt.Print("→ Masukan kode menu : ")
	fmt.Scan(&input)

	for input != "0" {

		// Mengecek ketersediaan menu
		if input == "1" {
			listAllTransaction(transactions)
			ShowListTransactionMenu()
		} else if input == "2" {
			listCustomerByDate(transactions)
			ShowListTransactionMenu()
		} else if input == "3" {
			listCustomerBySparepart(transactions, spareparts)
			ShowListTransactionMenu()
		} else if input == "4" {
			sFunc.ListBestSellerSparepart(&spareparts)
			ShowListTransactionMenu()
		} else {
			fmt.Println("Yah menu ga tersedia nih 😩")
		}

		// Meminta input dari user untuk memilih menu
		fmt.Print("→ Masukan kode menu : ")
		fmt.Scan(&input)
	}

	common.ResetConsole()
}

func listCustomerByDate(transactions tStruct.ArrTransaction) {
	common.ResetConsole()

	var startDateDay, startDateMonth, startDateYear int
	var endDateDay, endDateMonth, endDateYear int

	fmt.Println("=======================================================================================")
	fmt.Println("                       Lihat Pelanggan Berdasarkan Pembelian Periode                   ")
	fmt.Println("=======================================================================================")
	fmt.Println()

	fmt.Println("Masukan Tanggal Awal")
	fmt.Print("Contoh format 'Y-m-d': ")
	startDateDay, startDateMonth, startDateYear = common.InputDate()

	for startDateDay == -1 || startDateMonth == -1 || startDateYear == -1 {
		fmt.Println("Format tanggal tidak valid!")
		fmt.Println()

		fmt.Println("Masukan Tanggal Awal")
		fmt.Print("Contoh format 'Y-m-d': ")
		startDateDay, startDateMonth, startDateYear = common.InputDate()
	}

	fmt.Println("Masukan Tanggal Akhir")
	fmt.Print("Contoh format 'Y-m-d': ")
	endDateDay, endDateMonth, endDateYear = common.InputDate()

	for endDateDay == -1 || endDateMonth == -1 || endDateYear == -1 {
		fmt.Println("Format tanggal tidak valid!")
		fmt.Println()

		fmt.Println("Masukan Tanggal Akhir")
		fmt.Print("Contoh format 'Y-m-d': ")
		endDateDay, endDateMonth, endDateYear = common.InputDate()
	}

	common.ResetConsole()

	customers := GetTransactionCustomerByDate(transactions, startDateDay, startDateMonth, startDateYear, endDateDay, endDateMonth, endDateYear)
	fmt.Println("=======================================================================================")
	fmt.Println("                     List Pelanggan Yang Membeli")
	fmt.Println("=======================================================================================")
	fmt.Println()

	cFunc.ShowListCustomerPretty(customers)
	common.ShowEndAction(1)
	common.ResetConsole()
}

func listCustomerBySparepart(transactions tStruct.ArrTransaction, spareparts sStruct.ArrSparepart) {
	common.ResetConsole()

	var inputString string
	var sparepartIndex int = -1

	fmt.Println("=======================================================================================")
	fmt.Println("                     Lihat Pelanggan Berdasarkan Pembelian Sparepart                   ")
	fmt.Println("=======================================================================================")
	fmt.Println()
	fmt.Println("List Sparepart Tersedia: ")
	fmt.Println()
	sFunc.ShowSparepart(spareparts)
	fmt.Println()

	fmt.Print("Masukan ID Sparepart: ")

	for sparepartIndex == -1 {
		fmt.Scan(&inputString)

		sparepartIndex = sFunc.GetSparepartById(spareparts, inputString)
		if sparepartIndex == -1 {
			fmt.Print("ID tidak ditemukan, silahkan masukan kembali ID : ")
		}
	}

	common.ResetConsole()

	customers := GetTransactionCustomerBySparepart(transactions, spareparts.Data[sparepartIndex])

	fmt.Println("=======================================================================================")
	fmt.Println("                     List Pelanggan Yang Membeli", spareparts.Data[sparepartIndex].Name)
	fmt.Println("=======================================================================================")
	fmt.Println()

	cFunc.ShowListCustomerPretty(customers)
	common.ShowEndAction(1)
	common.ResetConsole()
}

func listAllTransaction(transactions tStruct.ArrTransaction) {
	common.ResetConsole()
	fmt.Println("=======================================================================================")
	fmt.Println("                                    List Transaksi                                     ")
	fmt.Println("=======================================================================================")
	fmt.Println()

	for i := 0; i < transactions.N; i++ {
		fmt.Print(i+1, ". ")
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
