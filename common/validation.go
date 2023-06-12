package common

/**
	Dibuat oleh Indra Mahesa 1302220067
**/
func CheckIfDateIsValid(date int) bool {
	if date >= 1 && date <= 31 {
		return true
	}
	return false
}

/**
	Dibuat oleh Indra Mahesa 1302220067
**/
func CheckIfMonthIsValid(month int) bool {
	if month >= 1 && month <= 12 {
		return true
	}
	return false
}