package piscine

func ListPushFront(l *List, data interface{}) {
	var new NodeL
	new.Data = data
	new.Next = (*NodeL)(nil)
	if (*l).Head == (*NodeL)(nil) && (*l).Tail == (*NodeL)(nil) {
		(*l).Head = &new
		(*l).Tail = &new
		return
	}
	new.Next = (*l).Head
	(*l).Head = &new
}

func ListSize(l *List) int {
	size := 0
	head := (*l).Head
	for head != nil {
		head = head.Next
		size++
	}
	return size
}
