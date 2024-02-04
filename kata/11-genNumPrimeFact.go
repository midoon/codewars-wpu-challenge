package kata

// https://www.codewars.com/kata/58f9f9f58b33d1b9cf00019d/train/go

// need improve

func CountFindNum(primes []int, limit int) []int {
	result := []int{}
	primesLArr := [][]int{}
	numArr := []int{}

	for i := 0; i < len(primes); i++ {
		num := 1
		primesL := primes
		isNotInitVal := false
		for {
			for _, val_j := range primesL {
				num = num * val_j
			}

			if num <= limit {
				if i == 0 {
					primesLArr = append(primesLArr, primesL)
					numArr = append(numArr, num)
				}

				if isNotInitVal && i != 0 {
					primesLArr = append(primesLArr, primesL)
					numArr = append(numArr, num)
				}
			} else {
				break
			}

			num = 1

			primesL = append(primesL, primes[i])

			isNotInitVal = true
		}

		primesL = primes

	}

	if len(numArr) == 0 {
		return []int{}
	}

	maxNum := numArr[0]

	for _, value := range numArr {
		if value > maxNum {
			maxNum = value
		}
	}

	result = append(result, len(primesLArr))
	result = append(result, maxNum)

	return result
}
