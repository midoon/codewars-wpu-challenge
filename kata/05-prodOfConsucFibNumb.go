package kata

// https://www.codewars.com/kata/5541f58a944b85ce6d00006a/train/go

func ProductFib(prod uint64) [3]uint64 {
	fibMap := make(map[int]uint64)
	num1 := uint64(0)
	num2 := uint64(0)
	result := uint64(0)
	f_n1 := uint64(0)
	f_n2 := uint64(1)
	f_n := uint64(0)
	for i := 0; i <= int(prod); i++ {

		fibMap[i] = f_n
		f_n = f_n1 + f_n2
		f_n2 = f_n1
		f_n1 = f_n

		if i > 1 {
			if fibMap[i]*fibMap[i-1] == prod {

				num1 = fibMap[i-1]
				num2 = fibMap[i]
				result = 1
				break
			} else if fibMap[i]*fibMap[i-1] > prod {
				num1 = fibMap[i-1]
				num2 = fibMap[i]
				result = 0
				break
			}
		}
	}
	return [3]uint64{uint64(num1), uint64(num2), uint64(result)}
}
