package transaction

import (
	"bengkelin/common"
	tStruct "bengkelin/features/transaction/structs"
	"fmt"
)

func EditTransaction(transactions *tStruct.ArrTransaction) {
	var input string

	common.ResetConsole()
	fmt.Println("====================================================================")
	fmt.Println("                         Edit Transaksi                             ")
	fmt.Println("====================================================================")
	fmt.Println("                         List Transaksi                             ")
	fmt.Println("                Pilih Transaksi ID yang ingin diubah                ")
	fmt.Println()

	for i := 0; i < transactions.N; i++ {
		fmt.Println(i+1, ".", "Transaksi ID : ", transactions.Data[i].Id)
		fmt.Printf("    Tanggal: %v-%v-%v ", transactions.Data[i].Year, transactions.Data[i].Month, transactions.Data[i].Date)
		fmt.Println()
		fmt.Printf("    Metode Pembayaran: %s ", transactions.Data[i].PaymentMethod)
		fmt.Println()
		fmt.Printf("    Harga: %d ", transactions.Data[i].Price)
		fmt.Println()
		fmt.Printf("    Catatan: %s ", transactions.Data[i].Note)
		fmt.Println()
		fmt.Printf("    Customer: %s ", transactions.Data[i].Customer.Name)
		fmt.Println()
	}

	fmt.Print("Ubah ID yang ingin diubah : ")
	fmt.Scan(&input)

	var i = GetTransactionById(*transactions, input)

	fmt.Println()
	DetailTransaction(transactions.Data[i])
	fmt.Println()

	if i == -1 {
		fmt.Println("GA ADA", input)
	}

	fmt.Print("Ubah Tanggal : ")
	fmt.Scan(&transactions.Data[i].Date)

	for !common.CheckIfDateIsValid(transactions.Data[i].Date) {
		fmt.Println("Tanggal tidak valid!")
		fmt.Println()
		fmt.Print("Ubah Tanggal : ")
		fmt.Scan(&transactions.Data[i].Date)
	}

	fmt.Print("Ubah Bulan : ")
	fmt.Scan(&transactions.Data[i].Month)

	for !common.CheckIfMonthIsValid(transactions.Data[i].Month) {
		fmt.Println("Bulan tidak valid!")
		fmt.Println()
		fmt.Print("Ubah Bulan : ")
		fmt.Scan(&transactions.Data[i].Month)
	}

	fmt.Print("Ubah Tahun : ")
	fmt.Scan(&transactions.Data[i].Month)

	fmt.Print("Ubah Metode Pembayaran : ")
	fmt.Scan(&transactions.Data[i].PaymentMethod)

	fmt.Print("Ubah Harga : ")
	fmt.Scan(&transactions.Data[i].Price)

	fmt.Print("Ubah Catatan : ")
	fmt.Scan(&transactions.Data[i].Note)

	fmt.Println("Berhasil Ubah Transaksi😊")
	common.ShowEndAction(1)

}
