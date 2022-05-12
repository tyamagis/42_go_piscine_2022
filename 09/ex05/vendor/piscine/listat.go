package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	for pos > 0 {
		l = (*l).Next
		if l == nil {
			return nil
		}
		pos--
	}
	return l
}
