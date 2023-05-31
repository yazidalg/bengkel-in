package sparepart

import (
	s "bengkelin/features/sparepart/structs"
	"fmt"
)

func AddSparepart(spareparts *s.ArrSparepart) {

	fmt.Println("0-----------------------------------------------0")
	fmt.Println("|		Tambah Data Sparepart		|")
	fmt.Println("0-----------------------------------------------0")
	inputSparepart(spareparts)
	printSparepart(*spareparts)
}

func inputSparepart(spareparts *s.ArrSparepart) {
	spareparts.N++
	fmt.Print("→ Masukan Nama : ")
	fmt.Scanf("%s", &spareparts.Data[spareparts.N-1].Name)
	fmt.Print("→ Masukan Harga : ")
	fmt.Scanf("%d", &spareparts.Data[spareparts.N-1].Harga)
	fmt.Print("→ Masukan Stok : ")
	fmt.Scanf("%d", &spareparts.Data[spareparts.N-1].Stok)
	spareparts.Data[spareparts.N].Id = spareparts.N

	fmt.Println("0--------------------------------------------0")
	fmt.Println("|    Yeay! Sparepart Berhasil Ditambahkan    |")
	fmt.Println("0--------------------------------------------0")
}

func printSparepart(spareparts s.ArrSparepart) {
	for i := 0; i < spareparts.N; i++ {
		fmt.Println(spareparts.Data[i].Name, spareparts.Data[i].Harga, spareparts.Data[i].Stok)
	}
}
