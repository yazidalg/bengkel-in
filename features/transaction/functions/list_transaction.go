package transaction

import (
	"bengkelin/common"
	cFunc "bengkelin/features/customer/functions"
	sFunc "bengkelin/features/sparepart/functions"
	sStruct "bengkelin/features/sparepart/structs"
	tStruct "bengkelin/features/transaction/structs"
	"fmt"
)

/*
*

	Dibuat oleh Indra Mahesa 1302220067

*
*/
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
			ListAllTransaction(transactions)
			common.ShowEndAction(1)
			common.ResetConsole()
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

/*
*

	Dibuat oleh Indra Mahesa 1302220067

*
*/
func listCustomerByDate(transactions tStruct.ArrTransaction) {
	common.ResetConsole()

	var day, month, year int

	fmt.Println("=======================================================================================")
	fmt.Println("                       Lihat Pelanggan Berdasarkan Pembelian Periode                   ")
	fmt.Println("=======================================================================================")
	fmt.Println()

	// Input Date
	fmt.Print("Masukan Tanggal : ")
	fmt.Scan(&day)

	for !common.CheckIfDateIsValid(day) {
		fmt.Println("Tanggal tidak valid!")
		fmt.Println()
		fmt.Print("Masukan Tanggal : ")
		fmt.Scan(&day)
	}

	// Input Month
	fmt.Print("Masukan Bulan : ")
	fmt.Scan(&month)

	for !common.CheckIfMonthIsValid(month) {
		fmt.Println("Bulan tidak valid!")
		fmt.Println()

		fmt.Print("Masukan Bulan : ")
		fmt.Scan(&month)
	}

	// Input Year
	fmt.Print("Masukan Tahun : ")
	fmt.Scan(&year)
	fmt.Println()

	common.ResetConsole()

	customers := GetTransactionCustomerByDate(transactions, day, month, year)

	fmt.Println("=======================================================================================")
	fmt.Println("                     List Pelanggan Yang Membeli")
	fmt.Println("=======================================================================================")
	fmt.Println()

	cFunc.ShowListCustomerPretty(customers)
	common.ShowEndAction(1)
	common.ResetConsole()
}

/*
*

	Dibuat oleh Indra Mahesa 1302220067

*
*/
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

/*
*

	Dibuat oleh Indra Mahesa 1302220067

*
*/
func ListAllTransaction(transactions tStruct.ArrTransaction) {
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
		fmt.Printf("   Metode Pembayaran: %v ", transactions.Data[i].PaymentMethod)
		fmt.Println()
		fmt.Printf("   Catatan: %v ", transactions.Data[i].Note)
		fmt.Println()
		fmt.Println()
		fmt.Println("- - - - - - - - - - - - - - - - ")
		fmt.Println()
	}
}
