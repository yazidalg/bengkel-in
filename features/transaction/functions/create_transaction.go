package transaction

import (
	"bengkelin/common"
	sStruct "bengkelin/features/sparepart/structs"
	tStruct "bengkelin/features/transaction/structs"
	"fmt"
)

func CreateTransaction(transactions *tStruct.ArrTransaction, spareparts *sStruct.ArrSparepart) bool {
	var inputInt int
	var inputString string

	
	fmt.Println("O-----------------------------------------------O")
	fmt.Println("|                 Tambah Transaksi              |")
	fmt.Println("|-----------------------------------------------|")
	fmt.Println("| List Sparepart tersedia :                       |")

	// Tampilkan Sparepart
	for i := 0; i < spareparts.N - 1; i++ {
		fmt.Print(spareparts.Data[i].Id, " - ", spareparts.Data[i].Name)
		fmt.Println()
	}

	// Input Date
	fmt.Print("| Masukan Tanggal : "); fmt.Scan(&inputInt)

	for !common.CheckIfDateIsValid(inputInt) {
		fmt.Println("| Tanggal tidak valid!                          |")
		fmt.Print("| Masukan Tanggal : "); fmt.Scan(&inputInt)
	}

	transactions.Data[0].Date = inputInt

	// Input Month
	fmt.Print("| Masukan Bulan : "); fmt.Scan(&inputInt)

	for !common.CheckIfMonthIsValid(inputInt) {
		fmt.Println("| Bulan tidak valid!                          |")
		fmt.Print("| Masukan Bulan : "); fmt.Scan(&inputInt)
	}

	transactions.Data[0].Month = inputInt

	// Input Year
	fmt.Print("| Masukan Tahun : "); fmt.Scan(&inputInt)
	transactions.Data[0].Year = inputInt

	// Input customerId
	fmt.Print("| Masukan ID Customer : "); fmt.Scan(&inputString)
	transactions.Data[0].CustomerId = inputString
	
	// Input paymentMethod
	fmt.Print("| Masukan Metode Pembayaran : "); fmt.Scan(&inputString)
	transactions.Data[0].PaymentMethod = inputString

	// Input paymentMethod
	fmt.Print("| Note : "); fmt.Scan(&inputString)
	transactions.Data[0].Note = inputString

	return false
}