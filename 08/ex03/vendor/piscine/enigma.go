package piscine

func isNil(a ***int, b *int, c *******int, d ****int) bool {
	if a == nil || *a == nil || **a == nil {
		return true
	}
	if b == nil {
		return true
	}
	if c == nil || *c == nil || **c == nil || ***c == nil || ****c == nil || *****c == nil || ******c == nil {
		return true
	}
	if d == nil || *d == nil || **d == nil || ***d == nil {
		return true
	}
	return false
}

func Enigma(a ***int, b *int, c *******int, d ****int) {
	if isNil(a, b, c, d) {
		return
	}
	***a, *b, *******c, ****d = *b, ****d, ***a, *******c
}
