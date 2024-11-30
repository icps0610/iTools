package utils

import (
	"fmt"
	"strconv"
)

func To_i(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func To_f(s string) float64 {
	fn, _ := strconv.ParseFloat(s, 64)
	return fn
}

func To_s(i int) string {
	return fmt.Sprintf(`%v`, i)
}
