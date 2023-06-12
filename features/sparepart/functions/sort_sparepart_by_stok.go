package sparepart

import (
	"bengkelin/common"
	s "bengkelin/features/sparepart/structs"
	"fmt"
)

/*
*

	Dibuat oleh Yazid Al Ghozali 1302223047

*
*/
func ListBestSellerSparepart(spareparts *s.ArrSparepart) {
	common.ResetConsole()

	var input string
	common.ShowSparepartSortMenu()
	fmt.Print("Masukan code : ")
	fmt.Scan(&input)

	for input != "0" {
		if input == "1" {
			common.ResetConsole()
			fmt.Println("=====================================================")
			fmt.Println("                  Pembelian Terbanyak                ")
			fmt.Println("=====================================================")
			fmt.Println()
			sortSparepartBySoldOut(spareparts, "DESC")
			printSparepartSoldOut(*spareparts)
			fmt.Println()
			common.ShowEndAction(1)
			common.ResetConsole()
		} else if input == "2" {
			common.ResetConsole()
			fmt.Println("=====================================================")
			fmt.Println("                   Pembelian Tedikit                 ")
			fmt.Println("=====================================================")
			fmt.Println()
			sortSparepartBySoldOut(spareparts, "ASC")
			printSparepartSoldOut(*spareparts)
			fmt.Println()
			common.ShowEndAction(1)
			common.ResetConsole()
		} else {
			fmt.Println("Yah menu ga tersedia nih 😩")
		}
		common.ShowSparepartSortMenu()
		fmt.Print("Masukan code : ")
		fmt.Scan(&input)
	}
	fmt.Println("Berhasil keluar🥳")
	common.ResetConsole()
}

/*
*

	Dibuat oleh Yazid Al Ghozali 1302223047

*
*/
func sortSparepartBySoldOut(spareparts *s.ArrSparepart, sort string) {
	var i, pass int
	var temp s.Sparepart
	pass = 1
	for pass < spareparts.N {
		i = pass
		temp = spareparts.Data[i]
		if sort == "ASC" {
			for i > 0 && temp.Sold_out < spareparts.Data[i-1].Sold_out {
				spareparts.Data[i] = spareparts.Data[i-1]
				i--
			}
		} else if sort == "DESC" {
			for i > 0 && temp.Sold_out > spareparts.Data[i-1].Sold_out {
				spareparts.Data[i] = spareparts.Data[i-1]
				i--
			}
		}
		spareparts.Data[i] = temp
		pass++
	}
}

/*
*

	Dibuat oleh Yazid Al Ghozali 1302223047

*
*/
func printSparepartSoldOut(spareparts s.ArrSparepart) {
	for i := 0; i < spareparts.N; i++ {
		fmt.Println("Nama Sparepart :", spareparts.Data[i].Name, "dibeli sebanyak", spareparts.Data[i].Sold_out)
	}
}
