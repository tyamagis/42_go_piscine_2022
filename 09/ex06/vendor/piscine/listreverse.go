package piscine

func recReverse(a, next *NodeL) {
	if next.Next != nil {
		recReverse(next, next.Next)
	}
	next.Next = a
}

func ListReverse(l *List) {
	recReverse(l.Head, l.Head.Next)
	l.Head, l.Tail = l.Tail, l.Head
	l.Tail.Next = nil
}
