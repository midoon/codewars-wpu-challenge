package kata

import "fmt"

// https://www.codewars.com/kata/5946a0a64a2c5b596500019a/train/go

func SplitAndAdd(numbers []int, n int) []int {

	result := []int{}
	numbersCopy := make([]int, len(numbers))
	copy(numbersCopy, numbers)

	splitTop := int(len(numbersCopy) / 2)
	splitDown := len(numbersCopy) - splitTop

	if n == 0 {
		return numbers
	}

	// your code is here
	for N := 0; N < n; N++ {
		result = []int{}

		for i := 0; i < splitDown; i++ {
			if splitTop != splitDown && i == 0 {
				result = append(result, numbersCopy[splitDown-1])
			}

			tempAddVal := 0
			if splitTop != splitDown && i != 0 {
				fmt.Println("top : ", numbersCopy[i-1])
				fmt.Println("down : ", numbersCopy[i+splitTop])
				tempAddVal = numbersCopy[i-1] + numbersCopy[i+splitTop]
				result = append(result, tempAddVal)
			}

			if splitTop == splitDown {
				fmt.Println("top : ", numbersCopy[i])
				fmt.Println("down : ", numbersCopy[i+splitTop])
				tempAddVal = numbersCopy[i] + numbersCopy[i+splitTop]
				result = append(result, tempAddVal)
			}

			tempAddVal = 0

		}
		fmt.Println("============================")

		numbersCopy = make([]int, len(result))
		copy(numbersCopy, result)
		splitTop = int(len(numbersCopy) / 2)
		splitDown = len(numbersCopy) - splitTop

	}

	return result
}
