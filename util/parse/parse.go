package parse

import (
	"math"
	"strconv"
)

func Hex2Dec(vals ...string) float64 {
	var result float64
	for index, val := range vals {
		floatVal, _ := strconv.ParseFloat(val, 64)
		result += math.Pow(256, float64(len(vals)-index)-1) * floatVal
	}
	return result
}