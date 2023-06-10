package sparepart

import s "bengkelin/features/sparepart/structs"

func GetSparepartById(spareparts s.ArrSparepart, id string) int {
	index := -1
	for i := 0; i < spareparts.N; i++ {
		if id == spareparts.Data[i].Id {
			index = i
		}
	}
	return index
}
