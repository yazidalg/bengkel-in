package customer

import cStruct "bengkelin/features/customer/structs"

/**
	Dibuat oleh Indra Mahesa 1302220067
	Binary Search
**/
func GetCustomerById(customers cStruct.ArrCustomer, id string) int {
	var found int = -1
	var med int

	var kr int = 0
	var kn int = customers.N - 1

	for kr <= kn && found == -1 {
		med = (kr + kn) / 2
		if id < customers.Data[med].Id {
			kn = med - 1
		} else if id > customers.Data[med].Id {
			kr = med + 1
		} else{
			found = med
		}
	}

	return found
}
