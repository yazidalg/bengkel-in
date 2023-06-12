package transaction

import (
	cStruct "bengkelin/features/customer/structs"
	tStruct "bengkelin/features/transaction/structs"
)

/**
	Dibuat oleh Indra Mahesa 1302220067
**/
func GetTransactionCustomerByDate(
	transactions tStruct.ArrTransaction, 
	day, 
	month, 
	year int,
) cStruct.ArrCustomer {
	var filteredCustomers cStruct.ArrCustomer

	for i := 0; i < transactions.N; i++ {
		if transactions.Data[i].Date == day && transactions.Data[i].Month == month && transactions.Data[i].Year == year {
			filteredCustomers.Data[filteredCustomers.N] = transactions.Data[i].Customer
			filteredCustomers.N++
		}
	}

	return filteredCustomers
}