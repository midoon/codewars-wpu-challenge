package kata

// https://www.codewars.com/kata/54521e9ec8e60bc4de000d6c/train/go

func maxNum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MaximumSubarraySum(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	maxSum, currentSum := numbers[0], numbers[0]

	for _, num := range numbers[1:] {
		currentSum = maxNum(num, currentSum+num)
		maxSum = maxNum(maxSum, currentSum)
	}

	if maxSum < 0 {
		return 0
	}

	return maxSum
}

// func MaximumSubarraySum(numbers []int) int {

//   max, curr := 0, 0
// 	for _, e := range numbers {
// 		if curr += e; curr < 0 {
// 			curr = 0
// 		} else if curr > max {
// 			max = curr
// 		}
// 	}
// 	return max
// }
