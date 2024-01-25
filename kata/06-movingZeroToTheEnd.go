package kata

//  https://www.codewars.com/kata/52597aa56021e91c93000cb0/train/go

func MoveZeros(arr []int) []int {
	// TODO: Program me
	var result []int
	var tempZero []int
	for _, val := range arr {
		if val != 0 {
			result = append(result, val)
		} else {
			tempZero = append(tempZero, val)
		}
	}

	result = append(result, tempZero...)
	return result
}
