package kata

// https://www.codewars.com/kata/525c65e51bf619685c000059/train/go

// func Cakes(recipe, available map[string]int) int {
// 	cake := 0

// 	if len(recipe) > len(available) {
// 		return 0
// 	}

// 	newAvb := make(map[string]int)

// 	for key, value := range available {
// 		newAvb[key] = value
// 	}

// 	// cakeMap := make(map[string]int)

// 	for {
// 		for k_rcp := range recipe {
// 			if _, ok := newAvb[k_rcp]; !ok {
// 				return 0
// 			} else {

// 				if newAvb[k_rcp]-recipe[k_rcp] < 0 {
// 					return cake
// 				}
// 				newAvb[k_rcp] = newAvb[k_rcp] - recipe[k_rcp]
// 			}
// 		}
// 		cake++
// 	}
// 	return cake
// }

func Cakes(recipe, available map[string]int) int {
	for k, v := range recipe {
		available[k] -= v
		if 0 > available[k] {
			return 0
		}
	}
	return 1 + Cakes(recipe, available)
}

// func Cakes(recipe, available map[string]int) int {
// 	var least int = 1e9
// 	for material, need := range recipe {
// 		if available[material]/need < least {
// 			least = available[material] / need
// 		}
// 	}
// 	return least
// }
