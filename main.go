package main

import (
	"bengkelin/common"
	s "bengkelin/features/sparepart"
	sStruct "bengkelin/features/sparepart/structs"
	t "bengkelin/features/transaction"
	tStruct "bengkelin/features/transaction/structs"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	var transactions tStruct.ArrTransaction

	ClearConsole()
	t.Main(&transactions)
	return
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
			// transaction.Main()
			common.ShowMainMenu()
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

// Fungsi tambahan untuk memudahkan membersihkan console
func ClearConsole() {
	c := exec.Command("reset")
	c.Stdout = os.Stdout
	c.Run()
}