package kata

// https://www.codewars.com/kata/523f5d21c841566fde000009/train/go

func ArrayDiff(a, b []int) []int {
	// your code here

	uniqueMap := make(map[int]bool)
	result := []int{}

	for _, b_val := range b {
		uniqueMap[b_val] = true
	}

	for _, val_a := range a {
		if _, ok := uniqueMap[val_a]; !ok {
			result = append(result, val_a)
		}
	}
	return result
}
