package sparepart

import s "bengkelin/features/sparepart/structs"

func SortSparepartByStok(spareparts s.ArrSparepart) {
	// var i, pass int
	// pass = 1
	
	// var temp s.ArrSparepart

	// for pass < spareparts.N {
	// 	temp = spareparts.Data[i].Sold_out
	// 	i = pass
	// 	for i > 0 && temp > spareparts.Data[i-1].Sold_out {
	// 		spareparts.Data[i].Stok = spareparts.Data[i-1]
	// 		i--
	// 	}
	// 	spareparts.Data[i].Stok = temp
	// 	i++
	// }
}

func insertion() {
	var temp, j int
	arr := []int{3, 4, 2, 1, 5}

	i := 1
	for i <= len(arr) - 1 {
					temp = arr[i]
					j = i

					for j > 0 && temp > arr[j - 1] {
									arr[j] = arr[j - 1]
									j--
					}

					arr[j] = temp
					i++
	}
}
