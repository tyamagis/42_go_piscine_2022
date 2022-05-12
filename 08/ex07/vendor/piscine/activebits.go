package piscine

func ActiveBits(n int) int {
	ct := 0
	bit := 1
	max := int((^uint(0)) >> 1)
	min := max + 1
	for bit != min {
		if n&bit == bit {
			ct++
		}
		bit <<= 1
	}
	if n < 0 {
		ct++
	}
	return ct
}
