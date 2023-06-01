package customer

import (
	cStruct "bengkelin/features/customer/structs"
	"fmt"
)

func ShowCustomer(customer cStruct.ArrCustomer) {
	fmt.Println("=======================================================================================")
	fmt.Println("                                     List Customer                                     ")
	fmt.Println("=======================================================================================")
	fmt.Println()
	for i := 0; i < customer.N; i++ {
		fmt.Print(i + 1, ". ")
		fmt.Println(
			"ID:", customer.Data[i].Id,
			"- Nama:", customer.Data[i].Name,
			"- Telp:", customer.Data[i].Phone,
			"- Alamat:", customer.Data[i].Address)
	}
}
