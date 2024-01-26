package kata

import "fmt"

// https://www.codewars.com/kata/513e08acc600c94f01000001/train/go

func hexClamp(val int) int {
	if val < 0 {
		return 0
	} else if val > 255 {
		return 255
	}
	return val
}

func RGB(r, g, b int) string {
	// Your code here

	R := hexClamp(r)
	G := hexClamp(g)
	B := hexClamp(b)

	return fmt.Sprintf("%02X%02X%02X", R, G, B)
}
