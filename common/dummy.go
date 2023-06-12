package common

import (
	cStruct "bengkelin/features/customer/structs"
	sStruct "bengkelin/features/sparepart/structs"
	tStruct "bengkelin/features/transaction/structs"
)

func CreateDummySparepart(spareparts *sStruct.ArrSparepart) {
	spareparts.N = 5
	spareparts.Data[0] = sStruct.Sparepart{"BDPFR", "Ban Dunloop", 101000, 10, 0}
	spareparts.Data[1] = sStruct.Sparepart{"RMRCX", "Rantai Motor", 130000, 5, 0}
	spareparts.Data[2] = sStruct.Sparepart{"KRMSS", "Kampas Rem", 25400, 0, 0}
	spareparts.Data[3] = sStruct.Sparepart{"OMNFS", "Oli Mesin", 100000, 1, 0}
	spareparts.Data[4] = sStruct.Sparepart{"SKTAF", "Seal Karet", 5500, 193, 0}
}

func CreateDummyTransaction(transactions *tStruct.ArrTransaction) {
	transactions.N = 2

	spareparts1 := [5]sStruct.Sparepart{
		{Id: "BDPFR", Name: "Ban Dunloop", Harga: 101000, Stok: 8, Sold_out: 0},
		{Id: "RMRCX", Name: "Rantai Motor", Harga: 130000, Stok: 5, Sold_out: 0},
		{Id: "KRMSS", Name: "Kampas Rem", Harga: 25400, Stok: 3, Sold_out: 0},
	}

	spareparts2 := [5]sStruct.Sparepart{
		{Id: "BDPFR", Name: "Ban Dunloop", Harga: 101000, Stok: 8, Sold_out: 0},
		{Id: "SKTAF", Name: "Seal Karet", Harga: 5500, Stok: 3, Sold_out: 0},
		{Id: "OMNFS", Name: "Oli Mesin", Harga: 100000, Stok: 5, Sold_out: 0},
		{Id: "RMRCX", Name: "Rantai Motor", Harga: 130000, Stok: 2, Sold_out: 0},
		{Id: "KRMSS", Name: "Kampas Rem", Harga: 25400, Stok: 1, Sold_out: 0},
	}

	transactions.Data[0] = tStruct.Transaction{
		Id:            "SIFHC",
		Note:          "-",
		PaymentMethod: "QRIS",
		Price:         10000,
		Date:          30,
		Month:         01,
		Year:          2023,
		Customer:      cStruct.Customer{Id: "4", Name: "Indra Mahesa", Address: "Bandung", Phone: "08123456789"},
		Spareparts:    sStruct.ArrSparepart{Data: spareparts1, N: 3},
	}

	// 2023-02-01
	// 2023-02-02

	transactions.Data[1] = tStruct.Transaction{
		Id:            "RMRCX",
		Note:          "-",
		PaymentMethod: "QRIS",
		Price:         5000,
		Date:          31,
		Month:         01,
		Year:          2003,
		Customer:      cStruct.Customer{Id: "3", Name: "Yazid Al Ghozalid", Address: "Cikrang", Phone: "08123456789"},
		Spareparts:    sStruct.ArrSparepart{Data: spareparts2, N: 4},
	}
}

func CreateDummyCustomer(customers *cStruct.ArrCustomer) {
	customers.N = 5
	customers.Data[0] = cStruct.Customer{"ZADCB", "Indra Mahesa", "Bandung", "08123456789"}
	customers.Data[1] = cStruct.Customer{"JBFEU", "Yazid Al Ghozalid", "Cikarang", "08123456789"}
	customers.Data[2] = cStruct.Customer{"BKBCD", "Wildan Syukri Niam", "Purwokerto", "08123456789"}
	customers.Data[3] = cStruct.Customer{"ANJSK", "Reinhard Efraim Situmeang", "Padang", "08123456789"}
	customers.Data[4] = cStruct.Customer{"GNSJB", "Aaron Joseph", "Jakarta", "08123456789"}
}
