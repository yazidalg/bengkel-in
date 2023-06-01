package sparepart

import (
	sStruct "bengkelin/features/sparepart/structs"
	"fmt"
)

func DetailSparepart(sparepart *sStruct.Sparepart) {
	fmt.Println()

	fmt.Println("Id :", sparepart.Id)
	fmt.Println("Nama :", sparepart.Name)
	fmt.Println("Harga : Rp.", sparepart.Harga)
	fmt.Println("Stok :", sparepart.Stok)

	fmt.Println()
}
