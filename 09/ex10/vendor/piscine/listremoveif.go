package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	n := l.Head
	prev := l.Head
	for i := 0; n != nil; {
		if i == 0 && n.Data == data_ref {
			l.Head = n.Next
			prev = l.Head
			i++
		} else if n.Data == data_ref {
			prev.Next = n.Next
		} else {
			prev = n
		}
		n = n.Next
	}
}
