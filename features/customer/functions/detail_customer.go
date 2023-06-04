package customer

import (
	cStruct "bengkelin/features/customer/structs"
	"fmt"
)

func DetailCustomer(customer cStruct.Customer) {
	fmt.Println("ID :", customer.Id)
	fmt.Println("Nama :", customer.Name)
	fmt.Println("Telp :", customer.Phone)
	fmt.Println("Alamat : ", customer.Address)
}
