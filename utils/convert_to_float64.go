package utils

import "strconv"

// ConvertToFloat64 function is used to convert the column to number.
func ConvertToFloat64(data string) (bool, float64) {
	result, err := strconv.ParseFloat(data, 64)
	if err != nil {
		return false, 0
	}
	return true, result
}
