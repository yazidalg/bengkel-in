package transaction

import (
	sTrancasction "bengkelin/features/transaction/structs"
)

/**
	Dibuat oleh Indra Mahesa 1302220067
**/
func GetTransactionById(transactions sTrancasction.ArrTransaction, id string) int {
	index := -1
	i := 0

	for i < transactions.N && index == -1 {
		if id == transactions.Data[i].Id {
			index = i
		}
		i++
	}
	
	return index
}