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

			cFunc.CreateCustomer(customers)
			cFunc.ShowCustomerMenu()

		} else if input == "2" {

			common.ResetConsole()
			fmt.Println("=======================================================================================")
			fmt.Println("                                     List Customer                                     ")
			fmt.Println("=======================================================================================")
			fmt.Println()
			cFunc.ShowCustomer(*customers)

			common.ShowEndAction(1)
			common.ResetConsole()

			cFunc.ShowCustomerMenu()
			
		} else if input == "3" {

			cFunc.EditCustomer(customers)
			cFunc.ShowCustomerMenu()

		} else if input == "4" {
			cFunc.DeleteCustomer(customers)
			cFunc.ShowCustomerMenu()
		} else {
			fmt.Println("Yah menu ga tersedia nih 😩")
		}

		// Meminta input dari user untuk memilih menu
		fmt.Print("→ Masukan kode menu : ")
		fmt.Scan(&input)
	}
}