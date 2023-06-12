package sparepart

import (
	common "bengkelin/common"
	sStruct "bengkelin/features/sparepart/structs"
	"fmt"
)

/**
	Dibuat oleh Yazid Al Ghozali 1302223047
**/
func DeleteSparepart(sparepart *sStruct.ArrSparepart) {
	common.ResetConsole()

	var i int
	var inputUser string

	i = -1
	fmt.Println("============================================")
	fmt.Println("              Hapus Sparepart               ")
	fmt.Println("============================================")
	fmt.Println()
	ShowSparepart(*sparepart)
	fmt.Println()
	if sparepart.N == 0 {
		fmt.Println("Wah data yang ingin dihapus masih kosong 😰")
		fmt.Println("Coba tambah data terlebih dahulu yaa 😁")
		fmt.Println()
		fmt.Println()
	} else {
		fmt.Print("Hapus sparepart berdasarkan ID : ")

		for i == -1 {
			fmt.Scan(&inputUser)
			i = GetSparepartById(*sparepart, inputUser)

			if i == -1 {
				fmt.Print("ID tidak ditemukan, silahkan masukan kembali ID yang ingin dihapus : ")
			}
		}

		common.ResetConsole()
		fmt.Println("=======================================================")
		fmt.Println("            Anda yakin ingin menghapus data?           ")
		fmt.Println("=======================================================")

		fmt.Println()
		DetailSparepart(sparepart.Data[i])
		fmt.Println()

		if common.ShowConfirmationMessage() {
			common.ResetConsole()
			for i := i; i < sparepart.N-1; i++ {
				sparepart.Data[i] = sparepart.Data[i+1]
			}
			sparepart.N--

			fmt.Println("=======================================================")
			fmt.Println("            Yeay sparepart berhasil dihapus🥳           ")
			fmt.Println("=======================================================")

			common.ShowEndAction(1)
			common.ResetConsole()
		} else {
			common.ResetConsole()
		}

	}
}
