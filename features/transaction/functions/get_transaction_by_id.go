package transaction

import (
	sTrancasction "bengkelin/features/transaction/structs"
)

/**
	Dibuat oleh Indra Mahesa 1302220067
**/
func GetTransactionById(transactions sTrancasction.ArrTransaction, id string) int {
	index := -1
	for i := 0; i < transactions.N; i++ {
		if id == transactions.Data[i].Id {
			index = i
		}
	}
	return index
}