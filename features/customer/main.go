package customer

import (
	"bengkelin/common"
	cFunc "bengkelin/features/customer/functions"
	cStruct "bengkelin/features/customer/structs"
	"fmt"
)

func Main(customers *cStruct.ArrCustomer) {
	// Reset Console
	common.ResetConsole()

	var input string

	// Menampilkan main menu
	cFunc.ShowCustomerMenu()

	// Meminta input dari user untuk memilih menu
	fmt.Print("→ Masukan kode menu : ")
	fmt.Scan(&input)
	
	for input != "5" {
		
		// Mengecek ketersediaan menu
		if input == "1" {
			common.ResetConsole()
			
			if cFunc.CreateCustomer(customers) {
				common.ResetConsole()
				common.ShowSuccessMessage()
				common.ShowEndAction()
			} else {
				common.ResetConsole()
				common.ShowSuccessMessage()
				common.ShowEndAction()
			}
			
			cFunc.ShowCustomerMenu()
		} else if input == "2" {
			common.ResetConsole()
			cFunc.ShowCustomer(*customers)
			common.ShowEndAction()
			common.ResetConsole()
			cFunc.ShowCustomerMenu()
		} else if input == "3" {
			fmt.Println("Yah menu ga tersedia nih 😩")
		} else if input == "4" {
			fmt.Println("Yah menu ga tersedia nih 😩")
		} else {
			fmt.Println("Yah menu ga tersedia nih 😩")
		}

		// Meminta input dari user untuk memilih menu
		fmt.Print("→ Masukan kode menu : ")
		fmt.Scan(&input)
	}
}