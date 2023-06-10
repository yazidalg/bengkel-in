package sparepart

const NMAX = 5

type Sparepart struct {
	Id, Name                      string
	Harga, Stok, Sold_out int
}

type ArrSparepart struct {
	Data [NMAX]Sparepart
	N    int
}
