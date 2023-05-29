package sparepart

type Sparepart struct {
	nama                      string
	id, harga, stok, sold_out int
}

type ArrSparepart [5]Sparepart