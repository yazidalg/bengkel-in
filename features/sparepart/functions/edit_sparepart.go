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
	searchDataSparepart(spareparts)
}

func searchDataSparepart(spareparts *s.ArrSparepart) {
	fmt.Print("→ Masukan ID yang ingin dicari : ")
	fmt.Scan(&spareparts.Data[spareparts.N].Id)
	showDetailSparepart(*spareparts, spareparts.Data[spareparts.N].Id)
}

func showDetailSparepart(spareparts s.ArrSparepart, id int) {
	for i := 0; i < spareparts.N; i++ {
		if id == spareparts.Data[i].Id {
			fmt.Println(spareparts.Data[i].Name)
		}
	}
}
