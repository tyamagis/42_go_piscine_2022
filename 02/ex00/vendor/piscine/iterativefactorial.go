package piscine

func IterativeFractorial(nb int) int {
	if (nb < 0 || nb > 20){
		return 0
	} else if nb == 0{
		return 1
	}
	ret := 1
	for ; nb > 1; nb-- {
		ret *= nb
	}
	return ret
}
