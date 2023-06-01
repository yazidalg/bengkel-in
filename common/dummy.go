package common

import (
	cStruct "bengkelin/features/customer/structs"
	sStruct "bengkelin/features/sparepart/structs"
)

func CreateDummySparepart(spareparts *sStruct.ArrSparepart) {
	spareparts.N = 5
	spareparts.Data[0] = sStruct.Sparepart{"Ban Dunloop", 1, 101000, 10, 0}
	spareparts.Data[1] = sStruct.Sparepart{"Rantai Motor", 2, 130000, 5, 0}
	spareparts.Data[2] = sStruct.Sparepart{"Kampas Rem", 3, 25400, 42, 0}
	spareparts.Data[3] = sStruct.Sparepart{"Oli Mesin", 4, 100000, 30, 0}
	spareparts.Data[4] = sStruct.Sparepart{"Seal Karet", 5, 5500, 193, 0}
}

func CreateDummyCustomer(customers *cStruct.ArrCustomer) {
	customers.N = 5
	customers.Data[0] = cStruct.Customer{"1", "Indra Mahesa", "Bandung", "08123456789"}
	customers.Data[1] = cStruct.Customer{"2", "Yazid Al Ghozalid", "Cikarang", "08123456789"}
	customers.Data[2] = cStruct.Customer{"3", "Wildan Syukri Niam", "Purwokerto", "08123456789"}
	customers.Data[3] = cStruct.Customer{"4", "Reinhard Efraim Situmeang", "Padang", "08123456789"}
	customers.Data[4] = cStruct.Customer{"5", "Aaron Joseph", "Jakarta", "08123456789"}

	// Id, Nama, Alamat, Telp
}