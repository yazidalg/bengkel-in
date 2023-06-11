package transaction

import (
	cStruct "bengkelin/features/customer/structs"
	tStruct "bengkelin/features/transaction/structs"
	"fmt"
)

// Psuedocode
// if start_year < transaction_year < end_year:
// 		filtered_transactions.append(transaction)
// elif start_year == transaction_year and start_month < transaction_month:
// 		filtered_transactions.append(transaction)
// elif end_year == transaction_year and transaction_month < end_month:
// 		filtered_transactions.append(transaction)
// elif start_year == transaction_year and start_month == transaction_month and start_day <= transaction_day:
// 		filtered_transactions.append(transaction)
// elif end_year == transaction_year and end_month == transaction_month and transaction_day <= end_day:
// 		filtered_transactions.append(transaction)

func GetTransactionCustomerByDate(
	transactions tStruct.ArrTransaction, 
	startDateDay, 
	startDateMonth, 
	startDateYear, 
	endDateDay, 
	endDateMonth, 
	endDateYear int,
) cStruct.ArrCustomer {
	var filteredCustomers cStruct.ArrCustomer

	for i := 0; i < transactions.N; i++ {
		if transactions.Data[i].Year > startDateYear && transactions.Data[i].Year < endDateYear {
			fmt.Println("Masuk Ke -1")
			filteredCustomers.Data[filteredCustomers.N] = transactions.Data[i].Customer
			filteredCustomers.N++
			} else if startDateYear == transactions.Data[i].Year && startDateMonth < transactions.Data[i].Month {
			fmt.Println("Masuk Ke -2")
			filteredCustomers.Data[filteredCustomers.N] = transactions.Data[i].Customer
			filteredCustomers.N++
			} else if endDateYear == transactions.Data[i].Year && transactions.Data[i].Month < endDateMonth {
			fmt.Println("Masuk Ke -3")
			filteredCustomers.Data[filteredCustomers.N] = transactions.Data[i].Customer
			filteredCustomers.N++
			} else if startDateYear == transactions.Data[i].Year && startDateMonth == transactions.Data[i].Month && startDateDay <= transactions.Data[i].Date {
			fmt.Println("Masuk Ke -4")
			filteredCustomers.Data[filteredCustomers.N] = transactions.Data[i].Customer
			filteredCustomers.N++
			} else if endDateYear == transactions.Data[i].Year && endDateMonth == transactions.Data[i].Month && transactions.Data[i].Date <= endDateDay {
			fmt.Println("Masuk Ke -5")
			filteredCustomers.Data[filteredCustomers.N] = transactions.Data[i].Customer
			filteredCustomers.N++
		}
	}

	return filteredCustomers
}