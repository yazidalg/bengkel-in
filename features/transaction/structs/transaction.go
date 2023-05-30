package transaction

const NMAX = 100

type Transaction struct {
	Id, Note, PaymentMethod, CustomerId string
	Date, Month, Year int
}

type ArrTransaction struct {
	Data [NMAX]Transaction
	N int
}