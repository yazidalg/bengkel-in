package transaction

import sparepart "bengkelin/features/sparepart"

type Transaction struct {
	id, note, paymentMethod, customerId string
	date, month, year int
	spareparts sparepart.ArrSparepart
}

type ArrTransaction [5]Transaction