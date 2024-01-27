package kata

// https://www.codewars.com/kata/521c2db8ddc89b9b7a0000c1

// i give up for this case

func Snail(snaipMap [][]int) []int {
	result := []int{}

	for len(snaipMap) > 0 {
		// Move from left to right
		for _, val := range snaipMap[0] {
			result = append(result, val)
		}
		snaipMap = snaipMap[1:]

		// Move from top to bottom
		for i := 0; i < len(snaipMap); i++ {
			result = append(result, snaipMap[i][len(snaipMap[i])-1])
			snaipMap[i] = snaipMap[i][:len(snaipMap[i])-1]
		}

		// Move from right to left
		if len(snaipMap) > 0 {
			lastRow := snaipMap[len(snaipMap)-1]
			for i := len(lastRow) - 1; i >= 0; i-- {
				result = append(result, lastRow[i])
			}
			snaipMap = snaipMap[:len(snaipMap)-1]
		}

		// Move from bottom to top
		for i := len(snaipMap) - 1; i >= 0; i-- {
			result = append(result, snaipMap[i][0])
			snaipMap[i] = snaipMap[i][1:]
		}
	}

	return result
}

// expected output []int{1, 2, 3, 6, 9, 8, 7, 4, 5}
