package piscine

func RecursiveFractorial(nb int) int {
	if nb == 0 {
		return 1
	} else if (nb < 0 || nb > 20){
		return 0
	}
	if nb == 1 {
		return nb
	}
	nb *= RecursiveFractorial(nb - 1)
	return nb
}
