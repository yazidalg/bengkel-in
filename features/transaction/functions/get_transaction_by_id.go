package transaction

import (
	sTrancasction "bengkelin/features/transaction/structs"
)

func GetTransactionById(transactions sTrancasction.ArrTransaction, id string) int {
	index := -1
	for i := 0; i < transactions.N; i++ {
		if id == transactions.Data[i].Id {
			index = i
		}
	}
	return index
}