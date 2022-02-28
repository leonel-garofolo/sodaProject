package utilities

import (
	"strconv"
)

func ParseIntNoError(s string) int {
	f, _ := strconv.Atoi(s)
	return f
}
func ParseFloatNoError(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}
