package kata

import (
	"strconv"
)

// https://www.codewars.com/kata/526571aae218b8ee490006f4/train/go

func CountBits(n uint) int {
	num := int64(n)
	numBit := strconv.FormatInt(num, 2)
	result := 0

	for _, str := range numBit {
		if string(str) == "1" {
			result += 1
		}
	}
	return result
}
