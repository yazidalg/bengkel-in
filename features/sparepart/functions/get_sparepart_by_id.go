package sparepart

import s "bengkelin/features/sparepart/structs"

/**
	Dibuat oleh Yazid Al Ghozali 1302223047
**/
func GetSparepartById(spareparts s.ArrSparepart, id string) int {
	index := -1
	i := 0

	for i < spareparts.N && index == -1 {
		if id == spareparts.Data[i].Id {
			index = i
		}

		i++
	}
	return index
}
