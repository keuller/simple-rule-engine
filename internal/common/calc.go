package common

import (
	"strconv"
	"strings"
)

func GetValue(ctype string, value string) float64 {
	if strings.ToLower(ctype) == "percentage" {
		num := value[:len(value)-1]
		val, _ := strconv.ParseFloat(num, 64)
		return (val / 100)
	}
	val, _ := strconv.ParseFloat(value, 64)
	return val
}
