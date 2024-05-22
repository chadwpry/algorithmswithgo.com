package module01

import (
	"math"
)

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//	BaseToDec("E", 16) => 14
//	BaseToDec("1110", 2) => 14
func BaseToDec(value string, base int) int {
	valueLen := len(value)

	if valueLen == 0 {
		return 0
	}

	powerBase := int(math.Pow(float64(base), float64(valueLen-1)))
	columnValue := string(value[0])

	return (powerBase * strToDecMap[columnValue]) + BaseToDec(value[1:], base)
}
