package piscine

func IterativePower(nb, pow int) int {
	ans := 1;
	for ; pow > 0; pow-- {
		ans *= nb
	}
	return ans
}
