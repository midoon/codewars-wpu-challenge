package kata

import "unicode"

// https://www.codewars.com/kata/52bc74d4ac05d0945d00054e/train/go

// stress

func FirstNonRepeating(str string) string {
	//your code here
	result := ""
	strMap := make(map[rune]int)
	for _, val := range str {
		if _, ok := strMap[unicode.ToLower(val)]; !ok {
			strMap[unicode.ToLower(val)] = 1
		} else {
			strMap[unicode.ToLower(val)] = strMap[unicode.ToLower(val)] + 1
		}

	}

	for _, val := range str {
		if strMap[unicode.ToLower(val)] == 1 {
			result = string(val)
			break
		}
	}

	return result
}

// func FirstNonRepeating(str string) string {
//     for _, c := range str {
//         if strings.Count(strings.ToLower(str), strings.ToLower(string(c))) < 2 {
// 	          return string(c)
// 	      }
//     }
//     return ""
// }
