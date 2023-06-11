package transaction

import (
	cStruct "bengkelin/features/customer/structs"
	tStruct "bengkelin/features/transaction/structs"
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
			filteredCustomers.Data[filteredCustomers.N] = transactions.Data[i].Customer
			filteredCustomers.N++
		} else if startDateYear == transactions.Data[i].Year && startDateMonth < transactions.Data[i].Month {
			filteredCustomers.Data[filteredCustomers.N] = transactions.Data[i].Customer
			filteredCustomers.N++
		} else if endDateYear == transactions.Data[i].Year && transactions.Data[i].Month < endDateMonth {
			filteredCustomers.Data[filteredCustomers.N] = transactions.Data[i].Customer
			filteredCustomers.N++
		} else if startDateYear == transactions.Data[i].Year && startDateMonth == transactions.Data[i].Month && startDateDay <= transactions.Data[i].Date {
			filteredCustomers.Data[filteredCustomers.N] = transactions.Data[i].Customer
			filteredCustomers.N++
		} else if endDateYear == transactions.Data[i].Year && endDateMonth == transactions.Data[i].Month && transactions.Data[i].Date <= endDateDay {
			filteredCustomers.Data[filteredCustomers.N] = transactions.Data[i].Customer
			filteredCustomers.N++
		}
	}

	return filteredCustomers
}