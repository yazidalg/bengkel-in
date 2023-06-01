package sparepart

import s "bengkelin/features/sparepart/structs"

func GetSparepartById(spareparts s.ArrSparepart, id int) int {
	index := -1
<<<<<<< HEAD

=======
>>>>>>> 8ff9c91 (onprogress: working on edit sparepart)
	for i := 0; i < spareparts.N; i++ {
		if id == spareparts.Data[i].Id {
			index = i
		}
	}
<<<<<<< HEAD

=======
>>>>>>> 8ff9c91 (onprogress: working on edit sparepart)
	return index
}
