package kata

import (
	"strings"
)

// https://www.codewars.com/kata/57d5e850bfcdc545870000b7/train/go

func max(a, b, c int) int {
	if a >= b && a >= c {
		return a
	} else if b >= a && b >= c {
		return b
	}
	return c
}

func DeadAntCount(ants string) int {
	// Remove occurrences of "ant" from the input
	cleanedString := strings.ReplaceAll(ants, "ant", "")

	// Count the remaining "a", "n", and "t" characters
	countA := strings.Count(cleanedString, "a")
	countN := strings.Count(cleanedString, "n")
	countT := strings.Count(cleanedString, "t")

	// Return the maximum count among "a", "n", and "t" characters
	return max(countA, countN, countT)
}
