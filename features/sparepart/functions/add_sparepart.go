package sparepart

import (
	common "bengkelin/common"
	s "bengkelin/features/sparepart/structs"
	"fmt"
)

func AddSparepart(spareparts *s.ArrSparepart) {
	common.ResetConsole()

	fmt.Println("0-----------------------------------------------0")
	fmt.Println("|		Tambah Data Sparepart		|")
	fmt.Println("0-----------------------------------------------0")
	if spareparts.N < s.NMAX {
		inputSparepart(spareparts)
	} else {
		common.ShowFullData()
		common.ShowEndAction(1)
		common.ResetConsole()
	}
}

func inputSparepart(spareparts *s.ArrSparepart) {
	var inputString string
	var inputInt int
	n := spareparts.N + 1
	i := spareparts.N
	fmt.Print("→ Masukan Nama : ")
	common.InputMultipleString(&inputString)
	spareparts.Data[i].Name = inputString
	fmt.Print("→ Masukan Harga : ")
	fmt.Scan(&inputInt)
	spareparts.Data[i].Harga = inputInt
	fmt.Print("→ Masukan Stok : ")
	fmt.Scan(&inputInt)
	spareparts.Data[i].Stok = inputInt
	spareparts.Data[spareparts.N].Id = common.GenerateRandomString(5)
	spareparts.N = n

	common.ResetConsole()

	fmt.Println("0--------------------------------------------0")
	fmt.Println("|    Yeay! Sparepart Berhasil Ditambahkan    |")
	fmt.Println("0--------------------------------------------0")

	common.ShowEndAction(1)
	common.ResetConsole()
}
