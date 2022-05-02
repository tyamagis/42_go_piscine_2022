package piscine

func MakeRange(min, max int) []int {
	len := max - min
	if len < 0 {
		len = 0
	}
	sl := make([]int, len)
	for i := 0; min < max; i, min = i+1, min+1 {
		sl[i] = min
	}
	return sl
}
