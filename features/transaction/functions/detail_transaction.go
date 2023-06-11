package transaction

import (
	sFunc "bengkelin/features/sparepart/functions"
	sTrancasction "bengkelin/features/transaction/structs"
	"fmt"
)

func DetailTransaction(transactions sTrancasction.Transaction) {
		fmt.Println("- - - - - - - - - - - - - - - - - - - - - - - - - - - -")
		fmt.Println()
		fmt.Println("                       Nota Transaksi                                 ")
		fmt.Println("                  Service Motor Bengkelin                             ")
		fmt.Println("               Jln. telekomunikasi rt 03 rw 12                             ")
		fmt.Println()
		fmt.Println("- - - - - - - - - - - - - - - - - - - - - - - - - - - -")
		fmt.Println()
		fmt.Println()
		fmt.Printf(" Tanggal: %v-%v-%v ", transactions.Year, transactions.Month, transactions.Date)
		fmt.Println(" Customer:", transactions.Customer.Name)
		fmt.Println()
		fmt.Println("- - - - - - - - - - - - - - - - - - - - - - - - - - - -")
		fmt.Println()
		for i := 0; i < transactions.Spareparts.N; i++ {
			fmt.Println(" ", transactions.Spareparts.Data[i].Name)
			fmt.Println(" ", transactions.Spareparts.Data[i].Harga, "x", transactions.Spareparts.Data[i].Stok)
		}
		fmt.Println()
		fmt.Println("- - - - - - - - - - - - - - - - - - - - - - - - - - - -")
		fmt.Println()
		fmt.Println(" Harga Sparepart: Rp.", sFunc.GetTotalPrice(transactions.Spareparts))
		fmt.Println(" Harga Tarif: Rp.", transactions.Price)
		fmt.Println(" Subtotal: Rp.", sFunc.GetTotalPrice(transactions.Spareparts) + transactions.Price)
		fmt.Println()
		fmt.Println("- - - - - - - - - - - - - - - - - - - - - - - - - - - -")
		fmt.Println()
}
