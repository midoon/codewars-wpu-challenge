package kata

// https://www.codewars.com/kata/58b38256e51f1c2af0000081/train/go

func BestMatch(a, b []int) int {
	// your code here

	result := 0
	gapPoint := a[0] - b[0]
	zamScore := 0
	for i := 0; i < len(a); i++ {
		if a[i]-b[i] < gapPoint {
			result = i
			gapPoint = a[i] - b[i]
			zamScore = b[i]
		} else if a[i]-b[i] == gapPoint {
			if zamScore < b[i] {
				zamScore = b[i]
				result = i
			}
		}
	}
	return result
}

// func BestMatch(a, b []int) int {
//   var min, index, goals int
//   for i,v := range b {
//     if dif := a[i]-v; i==0 || dif<min || dif==min && v>goals {
//       min, index, goals = dif, i, v
//     }
//   }
//   return index
// }
