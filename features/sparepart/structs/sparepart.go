package sparepart

const NMAX = 5

type Sparepart struct {
	Name                      string
	Id, Harga, Stok, Sold_out int
}

type ArrSparepart struct {
	Data [NMAX]Sparepart
	N    int
}
