package customer

import "fmt"

func Main() {
	var input string

	ShowCustomerMenu()

	// Meminta input dari user untuk memilih menu
	fmt.Print("→ Masukan kode menu : ")
	fmt.Scan(&input)
	
	for input != "4" {
		
		// Mengecek ketersediaan menu
		if input == "1" {
			fmt.Println("Anda berada di menu 1")
		} else if input == "2" {
			fmt.Println("Anda berada di menu 2")
		} else {
			fmt.Println("Yah menu ga tersedia nih 😩")
		}

		// Meminta input dari user untuk memilih menu
		fmt.Print("→ Masukan kode menu : ")
		fmt.Scan(&input)
	}
}