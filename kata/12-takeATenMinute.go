package kata

// https://www.codewars.com/kata/54da539698b8a2ad76000228/train/go

func IsValidWalk(walk []rune) bool {
	positionVal := 0
	directionVal := make(map[string]int)
	directionVal["s"] = -1
	directionVal["n"] = 1
	directionVal["e"] = 4
	directionVal["w"] = -4

	if len(walk)%2 != 0 || len(walk) != 10 {
		return false
	}

	for _, w := range walk {
		positionVal += directionVal[string(w)]
	}

	return positionVal == 0
}

// func IsValidWalk(walk []rune) bool {
//   if len(walk) != 10 {
//     return false
//   }

//   x, y := 0, 0
//   for _, r := range walk {
//     switch r {
//     case 'n': y++
//     case 's': y--
//     case 'e': x++
//     case 'w': x--
//     }
//   }

//   if x == 0 && y == 0 {
//     return true
//   }

//   return false
// }
