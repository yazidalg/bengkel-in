package sparepart

import sStruct "bengkelin/features/sparepart/structs"

func GetTotalPrice(spareparts sStruct.ArrSparepart) int {
	var totalPrice int = 0

	for i := 0; i < spareparts.N; i++ {
		totalPrice = totalPrice + spareparts.Data[i].Harga
	}

	return totalPrice
}