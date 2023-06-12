package customer

import (
	"bengkelin/common"
	cStruct "bengkelin/features/customer/structs"
	"fmt"
)

/**
	Dibuat oleh Indra Mahesa 1302220067
**/
func CreateCustomer(customers *cStruct.ArrCustomer) {
	common.ResetConsole()

	if customers.N < cStruct.NMAX {
		var inputString string

		n := customers.N + 1
		i := customers.N

		fmt.Println("=======================================================================================")
		fmt.Println("                                    Tambah Pelanggan                                   ")
		fmt.Println("=======================================================================================")
		fmt.Println()

		// Input Name
		fmt.Print(" Masukan Nama Pelanggan : ")
		common.InputMultipleString(&inputString)
		customers.Data[i].Name = inputString

		// Input Phone
		fmt.Print(" Masukan No Telp : ")
		fmt.Scan(&inputString)
		customers.Data[i].Phone = inputString

		// Input Address
		fmt.Print(" Masukan No Alamat : ")
		common.InputMultipleString(&inputString)
		customers.Data[i].Address = inputString

		fmt.Println("=======================================================================================")

		customers.N = n
		customers.Data[i].Id = common.GenerateRandomString(5)

		common.ResetConsole()
		common.ShowSuccessMessage()
		common.ShowEndAction(1)
	} else {
		common.ShowFullData()
		common.ShowEndAction(1)
		common.ResetConsole()
	}
}
