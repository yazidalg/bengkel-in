package sparepart

import (
	s "bengkelin/features/sparepart/structs"
	"fmt"
)

func EditSparepart(spareparts *s.ArrSparepart) {
	fmt.Println("0--------------------------------------------------------------0")
	fmt.Println("|                                                              |")
	fmt.Println("|                    Edit Data Sparepart                       |")
	fmt.Println("|          Cari Data Sparepart Dahulu Untuk Diubah             |")
	fmt.Println("|             Cari Data Sparepart Berdasarkan ID               |")
	fmt.Println("|                         Mwah😘                               |")
	fmt.Println("|                                                              |")
	fmt.Println("0--------------------------------------------------------------0")
	searchDataSparepartById(spareparts)
}

func searchDataSparepartById(spareparts *s.ArrSparepart) {
	fmt.Print("→ Masukan ID Sparepart yang akan dirubah : ")
	fmt.Scan(&spareparts.Data[spareparts.N].Id)
	showDetailSparepartById(*spareparts, spareparts.Data[spareparts.N].Id)
}

func showDetailSparepartById(spareparts s.ArrSparepart, id int) {
	for i := 0; i < spareparts.N; i++ {
		if id == spareparts.Data[i].Id {
			fmt.Println("0----------------------------------------0")
			fmt.Println("|                                        |")
			fmt.Println("| 1.", spareparts.Data[i].Name, "         ")
			fmt.Println("| 2.", spareparts.Data[i].Harga, "        ")
			fmt.Println("| 3.", spareparts.Data[i].Stok, "         ")
			fmt.Println("|                                        |")
			fmt.Println("0----------------------------------------0")
		}
	}
	if id >= spareparts.N {
		fmt.Println("0----------------------------------------0")
		fmt.Println("|                                        |")
		fmt.Println("|       Yah data tidak ditemukan😥       |")
		fmt.Println("|                                        |")
		fmt.Println("0----------------------------------------0")
	}
}
