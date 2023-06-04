package customer

import cStruct "bengkelin/features/customer/structs"

func GetCustomerById(customers cStruct.ArrCustomer, id string) int {
	index := -1
	for i := 0; i < customers.N; i++ {
		if id == customers.Data[i].Id {
			index = i
		}
	}
	return index
}
