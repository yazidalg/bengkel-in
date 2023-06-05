package sparepart

import (
	common "bengkelin/common"
	sStruct "bengkelin/features/sparepart/structs"
	"fmt"
)

func DeleteSparepart(sparepart *sStruct.ArrSparepart) {
	common.ResetConsole()

	var inputUser, i int
	i = -1
	fmt.Println("============================================")
	fmt.Println("              Hapus Sparepart               ")
	fmt.Println("============================================")
	ShowSparepart(*sparepart)
	fmt.Println()
	fmt.Println("Hapus sparepart berdasarkan ID : ")

	for i == -1 {
		fmt.Scan(&inputUser)
		i = GetSparepartById(*sparepart, inputUser)

		if i == -1 {
			fmt.Println("ID tidak ditemukan, silahkan masukan kembali ID yang ingin dihapus : ")
		}
	}

	DetailSparepart(&sparepart.Data[i])

	for i := i; i < sparepart.N-1; i++ {
		sparepart.Data[i] = sparepart.Data[i+1]
	}
	sparepart.N--

}
