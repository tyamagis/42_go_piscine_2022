package piscine

func IterativePower(nb, pow int) int {
	if pow < 0 {
		return 0
	}
	ans := 1;
	for ; pow > 0; pow-- {
		ans *= nb
	}
	return ans
}
