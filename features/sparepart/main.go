package sparepart

import (
	"bengkelin/common"
	sparepart "bengkelin/features/sparepart/functions"
	sStruct "bengkelin/features/sparepart/structs"
	"fmt"
)

func Main(spareparts *sStruct.ArrSparepart) {
	var input string

	// Menampilkan main menu
	sparepart.ShowSparepartMenu()

	// Meminta input dari user untuk memilih menu
	fmt.Print("→ Masukan kode menu : ")
	fmt.Scan(&input)

	for input != "5" {

		// Mengecek ketersediaan menu
		if input == "1" {
			sparepart.AddSparepart(spareparts)
			sparepart.ShowSparepartMenu()
		} else if input == "2" {
			common.ResetConsole()
			sparepart.ShowSparepart(*spareparts)
			common.ShowEndAction()
			sparepart.ShowSparepartMenu()
		} else if input == "3" {
			sparepart.EditSparepart(spareparts)
			sparepart.ShowSparepartMenu()
		} else if input == "4" {
			fmt.Println("Anda berada di menu 4")
		} else {
			fmt.Println("Yah menu ga tersedia nih 😩")
		}

		// Meminta input dari user untuk memilih menu
		fmt.Print("→ Masukan kode menu : ")
		fmt.Scan(&input)
	}

	fmt.Println("Keluar dari menu sparepart 🥳")
}


panjang data : 5

data 0: Ban mobil ID 1
data 1: oli ID 3
data 2: Spion ID 4
data 3: Spekboard ID 5



kampas rem ID 2
