package piscine

func ListSize(l *List) int {
	size := 0
	head := (*l).Head
	for head != nil {
		head = head.Next
		size++
	}
	return size
}
