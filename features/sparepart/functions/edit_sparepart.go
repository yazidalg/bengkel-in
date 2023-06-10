package sparepart

import (
	"bengkelin/common"
	sStruct "bengkelin/features/sparepart/structs"
	"fmt"
)

func EditSparepart(spareparts *sStruct.ArrSparepart) {
	common.ResetConsole()

	var inputString string

	fmt.Println("=======================================================================================")
	fmt.Println("                                    Edit Sparepart                                     ")
	fmt.Println("=======================================================================================")
	fmt.Println()
	fmt.Print("Cari Sparepart Berdasarkan ID : ")
	fmt.Scan(&inputString)

	var i = GetSparepartById(*spareparts, inputString)

	fmt.Println()
	DetailSparepart(spareparts.Data[i])
	fmt.Println()

	if i == -1 {
		fmt.Println("GA ADA", inputString)
	}

	fmt.Print("→ Masukan Nama : ")
	fmt.Scan(&spareparts.Data[i].Name)
	fmt.Print("→ Masukan Harga : ")
	fmt.Scan(&spareparts.Data[i].Harga)
	fmt.Print("→ Masukan Stok : ")
	fmt.Scan(&spareparts.Data[i].Stok)

	common.ResetConsole()

	fmt.Println("0--------------------------------------------0")
	fmt.Println("|    Yeay! Sparepart Berhasil Mengubah Data  |")
	fmt.Println("0--------------------------------------------0")

	common.ShowEndAction(1)
}
