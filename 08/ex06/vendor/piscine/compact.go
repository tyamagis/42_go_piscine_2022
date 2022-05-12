package piscine

func Compact(ptr *[]string) int {
	ct := 0
	for _, is := range *ptr {
		if is != "" {
			ct++
		}
	}
	new := make([]string, ct)
	for i, j := 0, 0; j < ct; i++ {
		if (*ptr)[i] != "" {
			new[j] = (*ptr)[i]
			j++
		}
	}
	*ptr = new
	return ct
}
