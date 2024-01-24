package kata

// How can you tell an extrovert from an introvert at NSA?
// Va gur ryringbef, gur rkgebireg ybbxf ng gur BGURE thl'f fubrf.

// I found this joke on USENET, but the punchline is scrambled. Maybe you can decipher it?
// According to Wikipedia, ROT13 is frequently used to obfuscate jokes on USENET.

// For this task you're only supposed to substitute characters. Not spaces, punctuation, numbers, etc.

// https://www.codewars.com/kata/52223df9e8f98c7aa7000062/train/go

// example

// "EBG13 rknzcyr." -> "ROT13 example."

// "This is my first ROT13 excercise!" -> "Guvf vf zl svefg EBG13 rkprepvfr!"

// A=65 M=77 N=78 Z=90 || a=97 m=109 n=110 z=122

func Rot13(msg string) string {
	var byteSlice []byte
	// Your code here
	for _, char := range msg {

		if char >= 65 && char <= 77 {
			byteSlice = append(byteSlice, byte(char+13))
		} else if char >= 78 && char <= 90 {
			byteSlice = append(byteSlice, byte(char-13))
		} else if char >= 97 && char <= 109 {
			byteSlice = append(byteSlice, byte(char+13))
		} else if char >= 110 && char <= 122 {
			byteSlice = append(byteSlice, byte(char-13))
		} else {
			byteSlice = append(byteSlice, byte(char))
		}
	}
	return string(byteSlice)
}
