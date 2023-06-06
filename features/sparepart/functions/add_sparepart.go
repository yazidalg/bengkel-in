package sparepart

import (
	common "bengkelin/common"
	s "bengkelin/features/sparepart/structs"
	"fmt"
)

func AddSparepart(spareparts *s.ArrSparepart) {

	fmt.Println("0-----------------------------------------------0")
	fmt.Println("|		Tambah Data Sparepart		|")
	fmt.Println("0-----------------------------------------------0")
	if spareparts.N < s.NMAX {
		inputSparepart(spareparts)
	} else {
		common.ShowFullData()
	}
}

func inputSparepart(spareparts *s.ArrSparepart) {
	n := spareparts.N + 1
	i := spareparts.N
	fmt.Print("→ Masukan Nama : ")
	fmt.Scan(&spareparts.Data[i].Name)
	fmt.Print("→ Masukan Harga : ")
	fmt.Scan(&spareparts.Data[i].Harga)
	fmt.Print("→ Masukan Stok : ")
	fmt.Scan(&spareparts.Data[i].Stok)
	spareparts.Data[spareparts.N].Id = common.GenerateRandomInt()
	spareparts.N = n

	fmt.Println("0--------------------------------------------0")
	fmt.Println("|    Yeay! Sparepart Berhasil Ditambahkan    |")
	fmt.Println("0--------------------------------------------0")
}
