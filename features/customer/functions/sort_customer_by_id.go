package customer

import (
	cStruct "bengkelin/features/customer/structs"
)

func SortCustomerById(customers cStruct.ArrCustomer) cStruct.ArrCustomer {

	var i, j, idx_min int
	var temp cStruct.Customer
	sortedCustomers := customers

	i = 1

	for i <= sortedCustomers.N-1 {
		idx_min = i - 1
		j = i
		for j < sortedCustomers.N {
			if sortedCustomers.Data[idx_min].Id > sortedCustomers.Data[j].Id {
				idx_min = j
			}
			j = j + 1
		}
		temp = sortedCustomers.Data[idx_min]
		sortedCustomers.Data[idx_min] = sortedCustomers.Data[i-1]
		sortedCustomers.Data[i-1] = temp
		i = i + 1
	}

	return sortedCustomers
}
