package transaction

import (
	"bengkelin/common"
	"fmt"
)

func CreateTransaction(transaction *ArrTransaction) bool {
	var inputInt int
	var inputString string

	
	fmt.Println("O-----------------------------------------------O")
	fmt.Println("|                 Tambah Transaksi              |")
	fmt.Println("|-----------------------------------------------|")
	fmt.Println("| List Sparepart tersedia :                       |")
	
	for spa

	// Input Date
	fmt.Print("| Masukan Tanggal : "); fmt.Scan(&inputInt)

	for !common.CheckIfDateIsValid(inputInt) {
		fmt.Println("| Tanggal tidak valid!                          |")
		fmt.Print("| Masukan Tanggal : "); fmt.Scan(&inputInt)
	}

	transaction[0].date = inputInt

	// Input Month
	fmt.Print("| Masukan Bulan : "); fmt.Scan(&inputInt)

	for !common.CheckIfMonthIsValid(inputInt) {
		fmt.Println("| Bulan tidak valid!                          |")
		fmt.Print("| Masukan Bulan : "); fmt.Scan(&inputInt)
	}

	transaction[0].month = inputInt

	// Input Year
	fmt.Print("| Masukan Tahun : "); fmt.Scan(&inputInt)
	transaction[0].year = inputInt

	// Input customerId
	fmt.Print("| Masukan ID Customer : "); fmt.Scan(&inputString)
	transaction[0].customerId = inputString
	
	// Input paymentMethod
	fmt.Print("| Masukan Metode Pembayaran : "); fmt.Scan(&inputString)
	transaction[0].paymentMethod = inputString

	// Input paymentMethod
	fmt.Print("| Note : "); fmt.Scan(&inputString)
	transaction[0].note = inputString

	return false
}