package sparepart

import (
	common "bengkelin/common"
	sStruct "bengkelin/features/sparepart/structs"
	"fmt"
)

func DeleteSparepart(sparepart *sStruct.ArrSparepart) {
	common.ResetConsole()

	var inputUser int
	fmt.Println("============================================")
	fmt.Println("              Hapus Sparepart               ")
	fmt.Println("============================================")
	ShowSparepart(*sparepart)
	fmt.Println()
	fmt.Println("Hapus sparepart berdasarkan ID : ")
	fmt.Scan(&inputUser)

	data := GetSparepartById(*sparepart, inputUser)

	DetailSparepart(&sparepart.Data[data])

	for i := data; i < sparepart.N-1; i++ {
		sparepart.Data[i] = sparepart.Data[i+1]
	}
	sparepart.N--

}

// N = [
// 	0 => {
// 		id : 1
// 		nama : yazid
// 	},
// 	1 => {
// 		id : 2
// 		nama : wildan
// 	},
// 	2 => {
// 		id : 3,
// 		nama : indra
// 	},
// 	3 => {
// 		id : 4,
// 		nama : reinhard
// 	}
// ]
