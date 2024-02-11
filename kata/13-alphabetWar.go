package kata

import "fmt"

//  https://www.codewars.com/kata/5938f5b606c3033f4700015a/train/go

func AlphabetWar(fight string) string {

	result := "Let's fight again!"

	left := make(map[string]int)
	right := make(map[string]int)

	left["w"] = 4
	left["p"] = 3
	left["b"] = 2
	left["s"] = 1

	right["m"] = 4
	right["q"] = 3
	right["d"] = 2
	right["z"] = 1

	leftPoint := 0
	rightPoint := 0

	afterAirstrike := false
	beforeAirstrikeChar := "_"

	for i, val := range fight {
		fmt.Println("index : ", i, " char : ", string(val))
		fmt.Println("real airstrike : ", afterAirstrike)
		fmt.Println("negasi airstrike : ", !afterAirstrike)
		if _, ok := left[string(val)]; ok && !afterAirstrike {
			fmt.Println("sisi kiri masuk : ", string(val))
			leftPoint += left[string(val)]
		} else if _, ok := right[string(val)]; ok && !afterAirstrike {
			fmt.Println("sisi kanan masuk : ", string(val))
			rightPoint += right[string(val)]
		}

		if afterAirstrike {
			fmt.Println("&& after air strike &&")
			afterAirstrike = false
		}

		if string(val) == "*" {
			fmt.Println("## ada air strike ##")
			afterAirstrike = true
			// delete point char before * char
			if _, ok := left[beforeAirstrikeChar]; ok {
				if leftPoint != 0 {
					leftPoint -= left[beforeAirstrikeChar]
				}

			} else if _, ok := right[beforeAirstrikeChar]; ok {
				if rightPoint != 0 {
					rightPoint -= right[beforeAirstrikeChar]
				}
			}
		}

		fmt.Println("kiri ", leftPoint)
		fmt.Println("kanan ", rightPoint)
		fmt.Println("===============================")
		beforeAirstrikeChar = string(val)
	}

	if leftPoint > rightPoint {
		result = "Left side wins!"
	} else if leftPoint < rightPoint {
		result = "Right side wins!"
	}
	return result
}

// problemnya di beforeAirestrike Char

// how to solve? bagaimana cara solve 1 char yang dihimpit 2 airstrike

//

// func AlphabetWar(fight string) string {
//   m := map[byte]int{'w':-4,'p':-3,'b':-2,'s':-1,'m':4,'q':3,'d':2,'z':1}
//   score := 0
//   b := []byte(fight)
//   for i:= range b {
//     score += m[ b[i] ]
//     if b[i] != '*' {continue}
//     if i-1>=0 { score -= m[b[i-1]];}
//     if i+1<len(b) && b[i+1]!='*' { b[i+1]=' ' }
//   }
//   if score<0 {return "Left side wins!"}
//   if score==0 {return "Let's fight again!"}
//   return "Right side wins!"
// }
