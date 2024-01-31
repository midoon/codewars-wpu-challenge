package kata

import (
	"strings"
)

// https://www.codewars.com/kata/5264d2b162488dc400000001/train/go

func SpinWords(str string) string {
	// words := strings.Fields(str)
	// tempRune := []rune{}
	// for indx := 0; indx < len(words); indx++ {
	// 	if len(words[indx]) > 4 {
	// 		tempRune = []rune(words[indx])
	// 		for i, j := 0, len(tempRune)-1; i < j; i, j = i+1, j-1 {
	// 			tempRune[i], tempRune[j] = tempRune[j], tempRune[i]
	// 		}

	// 		words[indx] = string(tempRune)
	// 		tempRune = []rune{}
	// 	}
	// }

	words := strings.Fields(str)
	for indx := 0; indx < len(words); indx++ {
		if len(words[indx]) > 4 {
			tempRune := []rune(words[indx])
			for i, j := 0, len(tempRune)-1; i < j; i, j = i+1, j-1 {
				tempRune[i], tempRune[j] = tempRune[j], tempRune[i]
			}

			words[indx] = string(tempRune)
		}
	}

	return strings.Join(words, " ")
} // SpinWords
