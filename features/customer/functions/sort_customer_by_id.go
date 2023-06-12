package customer

import (
	cStruct "bengkelin/features/customer/structs"
)

/**
	Dibuat oleh Indra Mahesa 1302220067
	Selection Sort by ASC or DESC
**/
func SortCustomerById(customers *cStruct.ArrCustomer, sort string) {
	var i, j, idx_min int
	var temp cStruct.Customer
	sortedCustomers := customers

	i = 1

	for i <= sortedCustomers.N-1 {
		idx_min = i - 1
		j = i
		for j < sortedCustomers.N {

			if sort == "ASC" {
				if sortedCustomers.Data[idx_min].Id > sortedCustomers.Data[j].Id {
					idx_min = j
				}
			} else if sort == "DESC" {
				if sortedCustomers.Data[idx_min].Id < sortedCustomers.Data[j].Id {
					idx_min = j
				}
			}

			j = j + 1
		}
		temp = sortedCustomers.Data[idx_min]
		sortedCustomers.Data[idx_min] = sortedCustomers.Data[i-1]
		sortedCustomers.Data[i-1] = temp
		i = i + 1
	}
}
