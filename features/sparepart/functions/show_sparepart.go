package sparepart

import (
	s "bengkelin/features/sparepart/structs"
	"fmt"
)

/**
	Dibuat oleh Yazid Al Ghozali 1302223047
**/
func ShowSparepart(spareparts s.ArrSparepart) {
	for i := 0; i < spareparts.N; i++ {
		fmt.Println(
			i+1,
			". ID:", spareparts.Data[i].Id,
			"- Nama:", spareparts.Data[i].Name,
			"- Stok:", spareparts.Data[i].Stok,
			"- Harga: Rp.", spareparts.Data[i].Harga)
	}
}
