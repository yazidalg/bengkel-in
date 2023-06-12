package common

import (
	"fmt"
	"strconv"
)

/**
	Dibuat oleh Indra Mahesa 1302220067
**/
func InputDate() (day int, month int, year int) {
	var inputDate string
	var inputDay, inputMonth, inputYear int
	var val int64
	var err error

	fmt.Scan(&inputDate)
	
	// Exctrating Date
	val, err = strconv.ParseInt(inputDate[0:4], 10, 64)
	if err != nil {
		return -1, -1, -1
	}

	inputYear = int(val)
	
	val, err = strconv.ParseInt(inputDate[5:7], 10, 64)
	if err != nil {
		return -1, -1, -1
	}

	inputMonth = int(val)

	if !CheckIfMonthIsValid(inputMonth) {
		return -1, -1, -1 
	}

	if len(inputDate) < 10 {
		val, err = strconv.ParseInt(inputDate[8:9], 10, 64)
	} else {
		val, err = strconv.ParseInt(inputDate[8:10], 10, 64)
	}

	if err != nil {
		return -1, -1, -1
	}

	inputDay = int(val)

	if !CheckIfDateIsValid(inputDay) {
		return -1, -1, -1 
	}

	return inputDay, inputMonth, inputYear
}