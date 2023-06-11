package customer

import (
	"bengkelin/common"
	cStruct "bengkelin/features/customer/structs"
	"fmt"
)

func EditCustomer(customers *cStruct.ArrCustomer) {
	common.ResetConsole()

	var inputString string
	i := -1

	fmt.Println("=======================================================================================")
	fmt.Println("                                    Edit Pelanggan                                     ")
	fmt.Println("=======================================================================================")
	fmt.Println()
	ShowCustomer(*customers)
	fmt.Println()
	fmt.Print("Cari Customer Berdasarkan ID : ")

	
	for i == -1 {
		fmt.Scan(&inputString)
		i = GetCustomerById(*customers, inputString)

		if i == -1 {
			fmt.Print("ID tidak ditemukan, silahkan masukan kembali ID : ")
		}
	}

	common.ResetConsole()

	fmt.Println("=======================================================================================")
	fmt.Println("                              Pelanggan Yang Akan Di Edit                              ")
	fmt.Println("=======================================================================================")
	fmt.Println()
	DetailCustomer(customers.Data[i])
	fmt.Println()
	// Input Name
	fmt.Print("Masukan Nama Pelanggan Baru : ")
	common.InputMultipleString(&inputString)
	customers.Data[i].Name = inputString

	// Input Phone
	fmt.Print("Masukan No Telp Baru : ")
	fmt.Scan(&inputString)
	customers.Data[i].Phone = inputString

	// Input Address
	fmt.Print("Masukan Alamat Baru : ")
	common.InputMultipleString(&inputString)
	customers.Data[i].Address = inputString

	common.ResetConsole()
	fmt.Println("=======================================================================================")
	fmt.Println("                              Pelanggan Berhasil Di Edit                              ")
	fmt.Println("=======================================================================================")
	
	common.ShowEndAction(1)
	common.ResetConsole()
}