package sparepart

import (
	"bengkelin/common"
	sStruct "bengkelin/features/sparepart/structs"
	"fmt"
)

func EditSparepart(spareparts *sStruct.ArrSparepart) {
	common.ResetConsole()

	var inputString string
	i := -1

	fmt.Println("=======================================================================================")
	fmt.Println("                                    Edit Sparepart                                     ")
	fmt.Println("=======================================================================================")
	fmt.Println()
	ShowSparepart(*spareparts)
	fmt.Println()
	fmt.Print("Cari Sparepart Berdasarkan ID : ")

	for i == -1 {
		fmt.Scan(&inputString)
		i = GetSparepartById(*spareparts, inputString)

		if i == -1 {
			fmt.Print("Sparepart tidak ditemukan. Masukan kembali ID Spearepart : ")
		}
	}

	fmt.Println()
	DetailSparepart(spareparts.Data[i])
	fmt.Println()

	fmt.Print("→ Masukan Nama : ")
	common.InputMultipleString(&spareparts.Data[i].Name)
	fmt.Print("→ Masukan Harga : ")
	fmt.Scan(&spareparts.Data[i].Harga)
	fmt.Print("→ Masukan Stok : ")
	fmt.Scan(&spareparts.Data[i].Stok)

	common.ResetConsole()

	fmt.Println("0--------------------------------------------0")
	fmt.Println("|    Yeay! Sparepart Berhasil Mengubah Data  |")
	fmt.Println("0--------------------------------------------0")

	common.ShowEndAction(1)
	common.ResetConsole()
}
