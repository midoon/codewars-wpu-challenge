package kata

import "strconv"

// https://www.codewars.com/kata/541c8630095125aba6000c00/train/go

func DigitalRoot(n int) int {
	stringNum := strconv.Itoa(n)
	strSlice := []string{}
	tempResult := 0

	for {
		for _, strNum := range stringNum {
			strSlice = append(strSlice, string(strNum))
		}

		for _, strVal := range strSlice {
			i, err := strconv.Atoi(strVal)
			if err != nil {
				return 0
			}
			tempResult += i
		}

		stringNum = strconv.Itoa(tempResult)
		tempResult = 0
		strSlice = []string{}

		if len(stringNum) == 1 {
			break
		}
	}

	number, err := strconv.Atoi(stringNum)
	if err != nil {
		return 0
	}
	return number
}

// func DigitalRoot(n int) int {
// 	return (n-1)%9 + 1
// }
