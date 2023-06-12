package transaction

import (
	"bengkelin/common"
	tStruct "bengkelin/features/transaction/structs"
	"fmt"
)

/**
	Dibuat oleh Yazid Al Ghozali 1302223047
**/
func EditTransaction(transactions *tStruct.ArrTransaction) {
	var input string
	i := -1
	common.ResetConsole()
	fmt.Println("====================================================================")
	fmt.Println("                         Edit Transaksi                             ")
	fmt.Println("====================================================================")
	fmt.Println("                         List Transaksi                             ")
	fmt.Println("                Pilih Transaksi ID yang ingin diubah                ")
	fmt.Println()

	ListAllTransaction(*transactions)

	fmt.Print("Ubah ID yang ingin diubah : ")

	for i == -1 {
		fmt.Scan(&input)
		i = GetTransactionById(*transactions, input)

		if i == -1 {
			fmt.Print("ID tidak ditemukan, masukan kembali ID : ")
		}
	}

	fmt.Println()
	DetailTransaction(transactions.Data[i])
	fmt.Println()

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
	common.InputMultipleString(&transactions.Data[i].PaymentMethod)

	fmt.Print("Ubah Harga : ")
	fmt.Scan(&transactions.Data[i].Price)

	fmt.Print("Ubah Catatan : ")
	common.InputMultipleString(&transactions.Data[i].Note)

	fmt.Println("Berhasil Ubah Transaksi😊")
	common.ShowEndAction(1)
	common.ResetConsole()
}
