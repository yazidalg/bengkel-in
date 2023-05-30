package customer

const NMAX = 100

type Customer struct {
	Id, Nama, Alamat, Telp string
}

type ArrCustomer struct {
	Data [NMAX]Customer
	N int
}