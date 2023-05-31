package transaction

import sStruct "bengkelin/features/sparepart/structs"

const NMAX = 100

type Transaction struct {
	Id, Note, PaymentMethod, CustomerId string
	Date, Month, Year int
	Spareparts sStruct.ArrSparepart
}

type ArrTransaction struct {
	Data [NMAX]Transaction
	N int
}