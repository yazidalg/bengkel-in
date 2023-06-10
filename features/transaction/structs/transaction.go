package transaction

import (
	cStruct "bengkelin/features/customer/structs"
	sStruct "bengkelin/features/sparepart/structs"
)

const NMAX = 100

type Transaction struct {
	Id, Note, PaymentMethod string
	Price, Date, Month, Year int
	Customer cStruct.Customer
	Spareparts sStruct.ArrSparepart
}

type ArrTransaction struct {
	Data [NMAX]Transaction
	N int
}