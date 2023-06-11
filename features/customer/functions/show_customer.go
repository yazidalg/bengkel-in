package customer

import (
	cStruct "bengkelin/features/customer/structs"
	"fmt"
)

func ShowCustomer(customer cStruct.ArrCustomer) {
	for i := 0; i < customer.N; i++ {
		fmt.Print(i + 1, ". ")
		fmt.Println(
			"ID:", customer.Data[i].Id,
			"- Nama:", customer.Data[i].Name,
			"- Telp:", customer.Data[i].Phone,
			"- Alamat:", customer.Data[i].Address)
	}
}

func ShowListCustomerPretty(customer cStruct.ArrCustomer) {
	for i := 0; i < customer.N; i++ {
			fmt.Print(i + 1, ". ")
			fmt.Print("ID: ", customer.Data[i].Id)
			fmt.Println()
			fmt.Printf("   Nama: %v ", customer.Data[i].Name)
			fmt.Println()
			fmt.Printf("   Telp: %v ", customer.Data[i].Phone)
			fmt.Println()
			fmt.Printf("   Alamat: %v ", customer.Data[i].Address)
			fmt.Println()
			fmt.Println()
			fmt.Println("- - - - - - - - - - - - - - - - ")
			fmt.Println()
	}
}
