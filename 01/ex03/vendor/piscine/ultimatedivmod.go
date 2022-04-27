package piscine

func UltimateDivMod(a, b *int) {
	var div, mod int
	div = *a / *b
	mod = *a % *b
	*a = div
	*b = mod
}
