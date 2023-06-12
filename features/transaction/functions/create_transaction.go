package transaction

import (
	"bengkelin/common"
	cFunc "bengkelin/features/customer/functions"
	cStruct "bengkelin/features/customer/structs"
	sFunc "bengkelin/features/sparepart/functions"
	sStruct "bengkelin/features/sparepart/structs"
	tStruct "bengkelin/features/transaction/structs"
	"fmt"
)

/**
	Dibuat oleh Indra Mahesa 1302220067
**/
func CreateTransaction(customers cStruct.ArrCustomer, transactions *tStruct.ArrTransaction, spareparts *sStruct.ArrSparepart) {
	common.ResetConsole()

	if transactions.N < tStruct.NMAX {
		var inputInt int
		var inputString string

		var newTransaction tStruct.Transaction
		var carts sStruct.ArrSparepart

		fmt.Println("=======================================================================================")
		fmt.Println("                                    Tambah Transaksi                                   ")
		fmt.Println("=======================================================================================")
		fmt.Println()

		// Input Date
		fmt.Print("Masukan Tanggal : ")
		fmt.Scan(&inputInt)

		for !common.CheckIfDateIsValid(inputInt) {
			fmt.Println("Tanggal tidak valid!")
			fmt.Println()
			fmt.Print("Masukan Tanggal : ")
			fmt.Scan(&inputInt)
		}

		newTransaction.Date = inputInt
		newTransaction.Id = common.GenerateRandomString(5)

		// Input Month
		fmt.Print("Masukan Bulan : ")
		fmt.Scan(&inputInt)

		for !common.CheckIfMonthIsValid(inputInt) {
			fmt.Println("Bulan tidak valid!")
			fmt.Println()

			fmt.Print("Masukan Bulan : ")
			fmt.Scan(&inputInt)
		}

		newTransaction.Month = inputInt

		// Input Year
		fmt.Print("Masukan Tahun : ")
		fmt.Scan(&inputInt)
		newTransaction.Year = inputInt
		fmt.Println()
		
		// Input Customer
		inputCustomer(customers, &newTransaction.Customer)

		// Input Price
		fmt.Print("Harga Service : ")
		fmt.Scan(&inputInt)
		newTransaction.Price = inputInt

		// Input paymentMethod
		fmt.Print("Masukan Metode Pembayaran : ")
		common.InputMultipleString(&inputString)
		newTransaction.PaymentMethod = inputString

		// Input paymentMethod
		fmt.Print("Note : ")
		common.InputMultipleString(&inputString)
		newTransaction.Note = inputString


		// Input Sparepart
		inputSpareparts(*spareparts, &carts)
		newTransaction.Spareparts = carts

		// Print Receipt
		DetailTransaction(newTransaction)

		if common.ShowConfirmationMessage() {
			transactions.Data[transactions.N] = newTransaction
			transactions.N++

			updateSparepartStock(newTransaction, spareparts)
		}

		common.ResetConsole()
		fmt.Println("=======================================================================================")
		fmt.Println("                            Yeay, Berhasil Menambahkan Transaksi!                   ")
		fmt.Println("=======================================================================================")
		fmt.Println()

		common.ShowEndAction(1)
	} else {
		common.ShowFullData()
		common.ShowEndAction(1)
	}

	common.ResetConsole()	
}

/**
	Dibuat oleh Indra Mahesa 1302220067
**/
func updateSparepartStock(transaction tStruct.Transaction, spareparts *sStruct.ArrSparepart) {
	for i := 0; i < transaction.Spareparts.N; i++ {
		sparepartIndex := sFunc.GetSparepartById(*spareparts, transaction.Spareparts.Data[i].Id)
		
		if sparepartIndex != -1 {
			spareparts.Data[sparepartIndex].Stok -= transaction.Spareparts.Data[i].Stok
			spareparts.Data[sparepartIndex].Sold_out += transaction.Spareparts.Data[i].Stok
		}
	}
}

/**
	Dibuat oleh Indra Mahesa 1302220067
**/
func inputCustomer(customers cStruct.ArrCustomer, customer *cStruct.Customer) {	
	var inputString string

	fmt.Println("List Customer Tersedia: ")
	fmt.Println()
	cFunc.ShowCustomer(customers)
	fmt.Println()
	fmt.Print("Masukan ID Customer : ")

	var customerIndex int = -1

	for customerIndex == -1 {
		fmt.Scan(&inputString)
		customerIndex = cFunc.GetCustomerById(customers, inputString)

		if customerIndex == -1 {
			fmt.Print("ID tidak ditemukan, silahkan masukan kembali ID : ")
		}
	}

	*customer = customers.Data[customerIndex]
}

/**
	Dibuat oleh Indra Mahesa 1302220067
**/
func inputSpareparts(spareparts sStruct.ArrSparepart, carts *sStruct.ArrSparepart) {
	fmt.Println("- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -")
	fmt.Println()

	var inputString string

	for inputString != "stop" {
		common.ResetConsole()

		fmt.Println("=======================================================================================")
		fmt.Println("                                 Keranjang Sparepart                                   ")
		fmt.Println("=======================================================================================")
		fmt.Println()
		fmt.Println("List Sparepart Tersedia: ")
		fmt.Println()
		sFunc.ShowSparepart(spareparts)
		fmt.Println()
		
		if carts.N > 0 {
			fmt.Println("Keranjang Sparepart: ")
			fmt.Println()
			sFunc.ShowSparepart(*carts)
			fmt.Println()
			
			fmt.Println("- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -")
			fmt.Println()
			
			fmt.Println("Subtotal: Rp. ", sFunc.GetTotalPrice(*carts))

			fmt.Println()
		}
		fmt.Println("- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -")
		fmt.Println()

		fmt.Println("Tambahkan sparepart yang akan dibeli")
		fmt.Println("ketik 'stop' untuk selesai menambahkan...")
		fmt.Println()
		fmt.Print("Masukan ID Sparepart: ")

		var sparepartIndex int = -1

		for sparepartIndex == -1 && inputString != "stop" {
			fmt.Scan(&inputString)

			if inputString != "stop" {
				sparepartIndex = sFunc.GetSparepartById(spareparts, inputString)
		
				if sparepartIndex == -1 {
					fmt.Print("ID tidak ditemukan, silahkan masukan kembali ID : ")
				}
			}
		}

		if inputString != "stop" {
			common.ResetConsole()
		
			// Show Added Sparepart
			fmt.Println("Sparepart Ditemukan:")
			fmt.Println()
			sFunc.DetailSparepart(spareparts.Data[sparepartIndex])
			fmt.Println()
	
			if common.ShowCustomConfirmationMessage("Tekan enter untuk tambahkan sparepart atau ketik 'batal' untuk membatalkan") {
				// Add Sparepart to cart
				// Check if sparepart already exists then increase the stock
				var duplicatedSparepartIndex int = sFunc.GetSparepartById(*carts, spareparts.Data[sparepartIndex].Id)
	
				if duplicatedSparepartIndex == -1 {

					// Check If Stock is Available
					if spareparts.Data[sparepartIndex].Stok <= 0 {
						common.ResetConsole()
						fmt.Println("Stock Tidak Mencukupi!")
						common.ShowEndAction(1)
					} else {
						newCart := spareparts.Data[sparepartIndex]
						newCart.Stok = 1
			
						carts.Data[carts.N] = newCart
						carts.N++
					}
				} else {
					if (spareparts.Data[sparepartIndex].Stok - carts.Data[duplicatedSparepartIndex].Stok) <= 0 {
						common.ResetConsole()
						fmt.Println("Stock Tidak Mencukupi!")
						common.ShowEndAction(1)
					} else {
						carts.Data[duplicatedSparepartIndex].Stok++
					}
				}
			}
		}
	}
}
