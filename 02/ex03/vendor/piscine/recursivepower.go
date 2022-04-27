package piscine

func RecursivePower(nb, pow int) int {
	if pow <= 0 {
		return 1
	} else if pow == 1 {
		return nb
	}
	nb *= RecursivePower(nb, pow - 1)
	return nb
}
