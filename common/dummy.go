package common

import sStruct "bengkelin/features/sparepart/structs"

func CreateDummySparepart(spareparts *sStruct.ArrSparepart) {
	spareparts.N = 5
	spareparts.Data[0] = sStruct.Sparepart{"Ban Dunloop", 1, 101000, 10, 0}
	spareparts.Data[1] = sStruct.Sparepart{"Rantai Motor", 2, 130000, 5, 0}
	spareparts.Data[2] = sStruct.Sparepart{"Kampas Rem", 3, 25400, 42, 0}
	spareparts.Data[3] = sStruct.Sparepart{"Oli Mesin", 4, 100000, 30, 0}
	spareparts.Data[4] = sStruct.Sparepart{"Seal Karet", 5, 5500, 193, 0}
}