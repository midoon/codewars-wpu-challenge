package kata

// https://www.codewars.com/kata/585d7d5adb20cf33cb000235/train/go
func FindUniq(arr []float32) float32 {
	// Do the magic
	result := arr[0]
	uniqMap := make(map[float32]int)
	for _, val := range arr {
		uniqMap[val] = uniqMap[val] + 1
	}

	for key := range uniqMap {
		if uniqMap[key] == 1 {
			result = key
		}
	}

	return result
}

// func FindUniq(arr []float32) float32 {
//   if arr[0] != arr[1] && arr[1] == arr[2] { return arr[0] }
//   for i, v := range arr[1:] {
//     if v != arr[i] { return v }
//   }
//   return 0
// }
