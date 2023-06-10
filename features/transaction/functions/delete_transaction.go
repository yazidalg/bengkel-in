package transaction

import (
	"bengkelin/common"
	tStruct "bengkelin/features/transaction/structs"
	"fmt"
)

func DeleteTransaction(transcations *tStruct.ArrTransaction) {
	var i int
	var input string
	i = -1
	fmt.Println("---------------------------------")
	fmt.Println("         Hapus Transaction       ")
	fmt.Println("---------------------------------")
	fmt.Println("List ID Transaction yang ada")
	for i := 0; i < transcations.N; i++ {
		fmt.Println(i, transcations.Data[i].Id, ". Tanggal : ", transcations.Data[i].Date, transcations.Data[i].Month, transcations.Data[i].Year)
	}
	if transcations.N == 0 {
		fmt.Println("Wah data yang ingin dihapus masih kosong 😰")
		fmt.Println("Coba tambah data terlebih dahulu yaa 😁")
		fmt.Println()
		fmt.Println()
	} else {
		fmt.Print("Masukan ID Transaction yang ingin dihapus : ")
		for i == -1 {
			fmt.Scan(&input)
			i = GetTransactionById(*transcations, input)

			if i == -1 {
				fmt.Println("ID tidak ditemukan, silahkan masukan kembali ID yang ingin dihapus : ")
			}
		}
		fmt.Println("=======================================================")
		fmt.Println("            Anda yakin ingin menghapus data?           ")
		fmt.Println("=======================================================")

		fmt.Println()
		DetailTransaction(transcations.Data[i])
		fmt.Println()

		if common.ShowConfirmationMessage() {
			for i := i; i < transcations.N-1; i++ {
				transcations.Data[i] = transcations.Data[i+1]
			}
			transcations.N--

			fmt.Println("=======================================================")
			fmt.Println("            Yeay sparepart berhasil dihapus🥳           ")
			fmt.Println("=======================================================")

			common.ShowEndAction(1)
		} else {
			common.ResetConsole()
		}
		
	}
}
