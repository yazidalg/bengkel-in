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

const NMAX = 5

func CreateTransaction(customers cStruct.ArrCustomer, transactions *tStruct.ArrTransaction, spareparts *sStruct.ArrSparepart) {
	common.ResetConsole()

	var inputInt int
	var inputString string

	n := transactions.N

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

	transactions.Data[n].Date = inputInt
	transactions.Data[n].Id = common.GenerateRandomString(5)

	// Input Month
	fmt.Print("Masukan Bulan : ")
	fmt.Scan(&inputInt)

	for !common.CheckIfMonthIsValid(inputInt) {
		fmt.Println("Bulan tidak valid!")
		fmt.Println()

		fmt.Print("Masukan Bulan : ")
		fmt.Scan(&inputInt)
	}

	transactions.Data[n].Month = inputInt

	// Input Year
	fmt.Print("Masukan Tahun : ")
	fmt.Scan(&inputInt)
	transactions.Data[n].Year = inputInt
	fmt.Println()
	
	// Input Customer
	inputCustomer(customers, &transactions.Data[n].Customer)

	// Input Price
	fmt.Print("Harga Service : ")
	fmt.Scan(&inputInt)
	transactions.Data[n].Price = inputInt

	// Input paymentMethod
	fmt.Print("Masukan Metode Pembayaran : ")
	fmt.Scan(&inputString)
	transactions.Data[n].PaymentMethod = inputString

	// Input paymentMethod
	fmt.Print("Note : ")
	fmt.Scan(&inputString)
	transactions.Data[n].Note = inputString

	fmt.Println()

	// Input Sparepart
	inputSpareparts(*spareparts, &carts)
	transactions.Data[n].Spareparts = carts

	// Print Receipt
	DetailTransaction(transactions.Data[n])

	if common.ShowConfirmationMessage() {
		transactions.N++
	}

	common.ResetConsole()
}

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

func inputSpareparts(spareparts sStruct.ArrSparepart, carts *sStruct.ArrSparepart) {
	fmt.Println("- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -")
	fmt.Println()

	var inputString string

	for inputString != "stop" {
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
					newCart := spareparts.Data[sparepartIndex]
					newCart.Stok = 1
		
					carts.Data[carts.N] = newCart
					carts.N++
				} else {
					carts.Data[duplicatedSparepartIndex].Stok++
				}
			}
		}
	

		common.ResetConsole()
	}
}
