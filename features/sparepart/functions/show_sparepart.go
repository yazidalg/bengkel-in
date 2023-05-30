package sparepart

import (
	s "bengkelin/features/sparepart/structs"
	"fmt"
)

func ShowSparepart(spareparts *s.ArrSparepart) {
	for i := 0; i < spareparts.N; i++ {
		fmt.Println(
			"| ID:", spareparts.Data[i].Id,
			"- Nama:", spareparts.Data[i].Name,
			"- Stok:", spareparts.Data[i].Stok,
			"- Harga: Rp.", spareparts.Data[i].Harga)
	}

	// | 1. Oli Motor (3) - Rp. 100000
	// | 2. Oli Motor (3) - Rp. 100000
	// | 3. Oli Motor (3) - Rp. 100000
	// | 4. Oli Motor (3) - Rp. 100000
}
