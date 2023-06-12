package sparepart

import (
	sStruct "bengkelin/features/sparepart/structs"
	"fmt"
)

/**
	Dibuat oleh Yazid Al Ghozali 1302223047
**/
func DetailSparepart(sparepart sStruct.Sparepart) {
	fmt.Println("Id :", sparepart.Id)
	fmt.Println("Nama :", sparepart.Name)
	fmt.Println("Harga : Rp.", sparepart.Harga)
	fmt.Println("Stok :", sparepart.Stok)
}
