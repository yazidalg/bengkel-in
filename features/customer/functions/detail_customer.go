package customer

import (
	cStruct "bengkelin/features/customer/structs"
	"fmt"
)

/**
	Dibuat oleh Indra Mahesa 1302220067
**/
func DetailCustomer(customer cStruct.Customer) {
	fmt.Println("ID :", customer.Id)
	fmt.Println("Nama :", customer.Name)
	fmt.Println("Telp :", customer.Phone)
	fmt.Println("Alamat : ", customer.Address)
}
