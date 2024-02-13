package kata

// https://www.codewars.com/kata/5679aa472b8f57fb8c000047/train/go

func FindEvenIndex(arr []int) int {
	left := 0
	right := 0
	result := -1

	for i := 0; i < len(arr); i++ {

		for j := 0; j < i; j++ {
			left += arr[j]
		}

		for k := i + 1; k < len(arr); k++ {
			right += arr[k]
		}

		if left == right {
			result = i
			return result
		}

		left = 0
		right = 0
	}
	return result
}
