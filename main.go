package main

import (
	"bengkelin/common"
	s "bengkelin/features/sparepart"
	sStruct "bengkelin/features/sparepart/structs"
	"fmt"
)

func main() {
	var input string
	var spareparts sStruct.ArrSparepart

	// Menampilkan main menu
	common.ShowMainMenu()

	// Meminta input dari user untuk memilih menu
	fmt.Print("→ Masukan kode menu : ")
	fmt.Scan(&input)

	for input != "4" {

		// Mengecek ketersediaan menu
		if input == "1" {
			fmt.Println("Anda berada di menu 1")
		} else if input == "2" {
			s.Main(&spareparts)
			common.ShowMainMenu()
		} else if input == "3" {
			fmt.Println("Anda berada di menu 3")
		} else {
			fmt.Println("Yah menu ga tersedia nih 😩")
		}

		// Meminta input dari user untuk memilih menu
		fmt.Print("→ Masukan kode menu : ")
		fmt.Scan(&input)
	}

	fmt.Println("Berhasil keluar aplikasi 🥳")
}
