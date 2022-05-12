package piscine

func ListLast(l *List) *NodeL {
	if l.Tail == nil {
		return nil
	}
	return l.Tail
}
