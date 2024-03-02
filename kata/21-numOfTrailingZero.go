package kata

import "fmt"

// https://www.codewars.com/kata/52f787eb172a8b4ae1000a34/train/go

func Zeros(n int) int {

	count := 0
	for i := 5; n/i >= 1; i *= 5 {
		count += n / i
		fmt.Println("i : ", n/i)
	}
	return count

	// if n == 0 {
	// 	return 0
	// }

	// result := big.NewInt(1)
	// for i := n; i > 0; i-- {
	// 	result.Mul(result, big.NewInt(int64(i)))
	// }
	// resultString := result.String()
	// resultChars := []rune(resultString)

	// zeros := 0

	// for i := len(resultChars) - 1; i >= 0; i-- {
	// 	if string(resultChars[i]) == "0" {
	// 		zeros++
	// 	}
	// 	if string(resultChars[i]) != "0" {
	// 		break
	// 	}
	// }
	// return zeros
}

// Menghitung Jumlah Nol: Pada setiap iterasi, kita menambahkan hasil dari n/i ke variabel count. Mengapa kita menggunakan n/i? Karena faktorial n adalah hasil kali semua bilangan bulat positif dari 1 hingga n, maka jumlah faktor 5 dalam faktorial n menentukan jumlah nol di belakang. Misalnya, jika n=10, maka faktorialnya adalah 10! = 10*9*8*7*6*5*4*3*2*1. Faktorial ini memiliki dua faktor 5 (dari 5 dan 10), yang berarti ada dua nol di belakang.
