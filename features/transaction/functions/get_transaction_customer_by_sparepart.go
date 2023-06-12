package transaction

import (
	cFunc "bengkelin/features/customer/functions"
	cStruct "bengkelin/features/customer/structs"
	sFunc "bengkelin/features/sparepart/functions"
	sStruct "bengkelin/features/sparepart/structs"
	tStruct "bengkelin/features/transaction/structs"
)

/**
	Dibuat oleh Indra Mahesa 1302220067
**/
func GetTransactionCustomerBySparepart(transactions tStruct.ArrTransaction, sparepart sStruct.Sparepart) cStruct.ArrCustomer {
	var filteredCustomers cStruct.ArrCustomer

	for i := 0; i < transactions.N; i++ {
		sparepartIndex := sFunc.GetSparepartById(transactions.Data[i].Spareparts, sparepart.Id)
		if sparepartIndex != -1 {
			customerIndex := cFunc.GetCustomerById(filteredCustomers, transactions.Data[i].Customer.Id)
			if customerIndex == -1 {
				filteredCustomers.Data[filteredCustomers.N] = transactions.Data[i].Customer
				filteredCustomers.N++
			}
		}
	}

	return filteredCustomers
}