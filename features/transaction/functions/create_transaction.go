package transaction

import (
	"bengkelin/common"
	sFunc "bengkelin/features/sparepart/functions"
	sStruct "bengkelin/features/sparepart/structs"
	tStruct "bengkelin/features/transaction/structs"
	"fmt"
)

const NMAX = 5

func CreateTransaction(transactions *tStruct.ArrTransaction, spareparts *sStruct.ArrSparepart) bool {
	var inputInt int
	var inputString string
	// var inputSparepart string

	n := transactions.N + 1
	// i := transactions.N

	fmt.Println("O-----------------------------------------------O")
	fmt.Println("|                 Tambah Transaksi              |")
	fmt.Println("|-----------------------------------------------|")
	fmt.Println("| List Sparepart tersedia :                     |")
	fmt.Println("|                                               |")
	sFunc.ShowSparepart(*spareparts)
	fmt.Println("|                                               |")

	// var sparepartsN int
	// for inputSparepart != "stop" {
	// 	fmt.Print("| Masukan ID Sparepart : "); fmt.Scan(&inputSparepart)

	// 	if inputSparepart != "stop" {
	// 		transactions.Data[n].Spareparts.Data[transactions.Data[n].Spareparts.N] = getSparepartById()
	// 		fmt.Print("| Masukan Jumlah Sparepart : "); fmt.Scan(&inputInt)
	// 		transactions.Data[n].Spareparts.Data[transactions.Data[n].Spareparts.N].Stock = inputInt
	// 		transactions.Data[n].Spareparts.N++
	// 	}
	// }

	// Input Date
	fmt.Print("| Masukan Tanggal : ")
	fmt.Scan(&inputInt)

	for !common.CheckIfDateIsValid(inputInt) {
		fmt.Println("| Tanggal tidak valid!                          |")
		fmt.Print("| Masukan Tanggal : ")
		fmt.Scan(&inputInt)
	}

	transactions.Data[n].Date = inputInt

	// Input Month
	fmt.Print("| Masukan Bulan : ")
	fmt.Scan(&inputInt)

	for !common.CheckIfMonthIsValid(inputInt) {
		fmt.Println("| Bulan tidak valid!                          |")
		fmt.Print("| Masukan Bulan : ")
		fmt.Scan(&inputInt)
	}

	transactions.Data[n].Month = inputInt

	// Input Year
	fmt.Print("| Masukan Tahun : ")
	fmt.Scan(&inputInt)
	transactions.Data[n].Year = inputInt

	// Input customerId
	fmt.Print("| Masukan ID Customer : ")
	fmt.Scan(&inputString)
	transactions.Data[n].CustomerId = inputString

	// Input paymentMethod
	fmt.Print("| Masukan Metode Pembayaran : ")
	fmt.Scan(&inputString)
	transactions.Data[n].PaymentMethod = inputString

	// Input paymentMethod
	fmt.Print("| Note : ")
	fmt.Scan(&inputString)
	transactions.Data[n].Note = inputString

	return false
}

func inputSpareparts(transactions *tStruct.ArrTransaction) {

}
