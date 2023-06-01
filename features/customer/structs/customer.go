package customer

const NMAX = 100

type Customer struct {
	Id, Name, Address, Phone string
}

type ArrCustomer struct {
	Data [NMAX]Customer
	N int
}