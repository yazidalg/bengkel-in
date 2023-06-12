package sparepart

import (
	"bengkelin/common"
	s "bengkelin/features/sparepart/structs"
	"fmt"
)

func ListBestSellerSparepart(spareparts *s.ArrSparepart) {
	common.ResetConsole()
	
	var input string
	fmt.Println("O---------------------------------------------------------O")
	fmt.Println("|          Urutkan Sparepart Paling Banyak Dibeli         |")
	fmt.Println("|---------------------------------------------------------|")
	fmt.Println("| 1. Pembelian Terbanyak                                  |")
	fmt.Println("| 2. Pembelian Terdikit                                   |")
	fmt.Println("| 0. Kembali                                              |")
	fmt.Println("O---------------------------------------------------------O")
	fmt.Print("Masukan code : ")
	fmt.Scan(&input)

	common.ResetConsole()

	for input != "0" {
		if input == "1" {
			common.ResetConsole()
			fmt.Println("=====================================================")
			fmt.Println("            Mengurutkan Dari Besar Ke Kecil          ")
			fmt.Println("=====================================================")
			fmt.Println()
			sortSparepartDescending(spareparts)
			printSparepartSoldOut(*spareparts)
			fmt.Println()
			common.ShowEndAction(1)
			common.ResetConsole()
		} else if input == "2" {
			common.ResetConsole()
			fmt.Println("=====================================================")
			fmt.Println("            Mengurutkan Dari Kecil Ke Besar          ")
			fmt.Println("=====================================================")
			fmt.Println()
			sortSparepartAscending(spareparts)
			printSparepartSoldOut(*spareparts)
			fmt.Println()
			common.ShowEndAction(1)
			common.ResetConsole()
		} else {
			fmt.Println("Yah menu ga tersedia nih 😩")
		}
		fmt.Print("Masukan code : ")
		fmt.Scan(&input)
	}
	fmt.Println("Berhasil keluar🥳")
	common.ResetConsole()
}

func sortSparepartAscending(spareparts *s.ArrSparepart) {
	var i, pass, temp int
	pass = 1
	for pass < spareparts.N {
		i = pass
		temp = spareparts.Data[i].Sold_out
		for i > 0 && temp > spareparts.Data[i-1].Sold_out {
			spareparts.Data[i].Sold_out = spareparts.Data[i-1].Sold_out
			i--
		}
		spareparts.Data[i].Sold_out = temp
		pass++
	}
}

func sortSparepartDescending(spareparts *s.ArrSparepart) {
	var i, pass, temp int
	pass = 1
	for pass < spareparts.N {
		i = pass
		temp = spareparts.Data[i].Sold_out
		for i > 0 && temp < spareparts.Data[i-1].Sold_out {
			spareparts.Data[i].Sold_out = spareparts.Data[i-1].Sold_out
			i--
		}
		spareparts.Data[i].Sold_out = temp
		pass++
	}
}

func printSparepartSoldOut(spareparts s.ArrSparepart) {
	for i := 0; i < spareparts.N; i++ {
		fmt.Println("Nama Sparepart :", spareparts.Data[i].Name, "dibeli sebanyak", spareparts.Data[i].Sold_out)
	}
}
