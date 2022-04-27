package piscine

func TableLen(t []int) int {
	len := 0
	for range t {
		len++
	}
	return len
}

func Swap(a, b *int) {
	*a, *b = *b, *a
}

func SortIntegerTable(table []int) {
	len := TableLen(table)
	for i := 0; i < len-1; i++ {
		for j := i + 1; j < len; j++ {
			if table[i] > table[j] {
				Swap(&table[i], &table[j])
			}
		}
	}
}
