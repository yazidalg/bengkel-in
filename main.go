package main

import (
	"bengkelin/common"
	sparepart "bengkelin/features/sparepart"
	transaction "bengkelin/features/transaction"
	"fmt"
	"os"
	"os/exec"
)

const NMAX = 5

type Sparepart struct {
	nama string
	id, harga, stok, sold_out int
}

type arrSparepart [NMAX]Sparepart

func main() {
	var spareparts sparepart.ArrSparepart
	var transactions transaction.ArrTransaction

	ClearConsole()
	transaction.Main(&transactions, &spareparts)
	return
	var input string
	
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
			fmt.Println("Anda berada di menu 2")
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