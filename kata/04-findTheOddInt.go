package kata

// https://www.codewars.com/kata/54da5a58ea159efa38000836/train/go

// Given an array of integers, find the one that appears an odd number of times.

// There will always be only one integer that appears an odd number of times.

// [7] should return 7, because it occurs 1 time (which is odd).
// [0] should return 0, because it occurs 1 time (which is odd).
// [1,1,2] should return 2, because it occurs 1 time (which is odd).
// [0,1,0,1,0] should return 0, because it occurs 3 times (which is odd).
// [1,2,2,3,3,3,4,3,3,3,2,2,1] should return 4, because it appears 1 time (which is odd).

func FindOdd(seq []int) int {
	uniqueMap := make(map[int]bool)
	uniqSlice := []int{}
	var tempOdd int
	var result int

	for _, value := range seq {
		if _, ok := uniqueMap[value]; !ok {
			uniqueMap[value] = true
			uniqSlice = append(uniqSlice, value)
		}
	}

	for _, uniq := range uniqSlice {
		for _, val := range seq {
			if uniq == val {
				tempOdd++
			}
		}
		if tempOdd%2 != 0 {
			result = uniq
			tempOdd = 0
		}
	}

	return result
}
