package piscine

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	} else if nb == 2 {
		return true
	}
	if nb&1 == 1 {
		for i := 3; i <= nb/i; i += 2 {
			if nb%i == 0 {
				return false
			}
		}
		return true
	} else {
		return false
	}
}

// func IsPrime(nb int) bool {
// 	if nb <= 1 {
// 		return false
// 	} else if nb == 2 {
// 		return true
// 	}
// 	if nb&1 == 1 {
// 		for i := 3; i <= nb/i; i += 2 {
// 			if nb%i == 0 {
// 				return false
// 			}
// 		}
// 		return true
// 	} else {
// 		if nb == 2 {
// 			return true
// 		}
// 		return false
// 	}
// }

// func IsPrime(nb int) bool {
// 	if nb <= 1 {
// 		return false
// 	}
// 	if nb&1 == 1 {
// 		for i := 1; i <= nb/i; i += 2 {
// 			if (i != 1) && (nb%i == 0) {
// 				return false
// 			}
// 		}
// 		return true
// 	} else {
// 		if nb == 2 {
// 			return true
// 		}
// 		return false
// 	}
// }
