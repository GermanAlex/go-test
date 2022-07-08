package example

import (
	"math"
	"sort"
)

// CountLetters подсчитывает количество символов в строке.
func CountLetters(s string, l byte) int {
	c := 0
	for _, b := range []byte(s) {
		if b == l {
			c++
		}
	}
	return c
}

func SomeMath(rad float64) float64 {
	return math.Sin(rad) * math.Cos(rad)
}

func SortInts(intSlice []int) []int {
	sort.Ints(intSlice)
	return intSlice
}

func NineSqrt(res float64) float64 {
	return math.Sqrt(9)
}
