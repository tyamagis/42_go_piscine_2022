package piscine

func ListMerge(l1, l2 *List) {
	n := ListLast(l1)
	n.Next = l2.Head
	l1.Tail = l2.Tail
	l2.Head = nil
	l2.Tail = nil
}
