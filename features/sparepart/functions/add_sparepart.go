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
	if spareparts.N < s.NMAX-1 {
		inputSparepart(spareparts)
	} else {
		common.ShowFullData()
	}
}

func inputSparepart(spareparts *s.ArrSparepart) {
	spareparts.N++
	fmt.Print("→ Masukan Nama : ")
	fmt.Scan(&spareparts.Data[spareparts.N-1].Name)
	fmt.Print("→ Masukan Harga : ")
	fmt.Scan(&spareparts.Data[spareparts.N-1].Harga)
	fmt.Print("→ Masukan Stok : ")
	fmt.Scan(&spareparts.Data[spareparts.N-1].Stok)
	spareparts.Data[spareparts.N].Id = spareparts.N

	fmt.Println("0--------------------------------------------0")
	fmt.Println("|    Yeay! Sparepart Berhasil Ditambahkan    |")
	fmt.Println("0--------------------------------------------0")
}
