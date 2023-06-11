package customer

import (
	"bengkelin/common"
	cStruct "bengkelin/features/customer/structs"
	"fmt"
)

func DeleteCustomer(customers *cStruct.ArrCustomer) {
	common.ResetConsole()

	var inputString string
	i := -1

	fmt.Println("=======================================================================================")
	fmt.Println("                                    Hapus Customer                                     ")
	fmt.Println("=======================================================================================")
	fmt.Println()
	ShowCustomer(*customers)
	fmt.Println()
	fmt.Print("Cari ID Customer Yang Akan Di Hapus : ")
	
	for i == -1 {
		fmt.Scan(&inputString)
		i = GetCustomerById(*customers, inputString)

		if i == -1 {
			fmt.Print("ID tidak ditemukan, silahkan masukan kembali ID : ")
		}
	}
	
	common.ResetConsole()
	fmt.Println("=======================================================================================")
	fmt.Println("                     Apakah Anda Yakin Ingin Menghapus Customer Ini?                   ")
	fmt.Println("=======================================================================================")
	fmt.Println()
	DetailCustomer(customers.Data[i])
	fmt.Println()

	if common.ShowConfirmationMessage() {

		for i := i; i < customers.N; i++ {
			customers.Data[i] = customers.Data[i+1]
		}

		customers.N--

		common.ResetConsole()
		fmt.Println("=======================================================================================")
		fmt.Println("                               Customer Berhasil Di Hapus                              ")
		fmt.Println("=======================================================================================")
		common.ShowEndAction(1)
	}	
	
	common.ResetConsole()
}