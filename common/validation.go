package common

func CheckIfDateIsValid(date int) bool {
	if date >= 1 && date <= 31 {
		return true
	}
	return false
}

func CheckIfMonthIsValid(month int) bool {
	if month >= 1 && month <= 12 {
		return true
	}
	return false
}