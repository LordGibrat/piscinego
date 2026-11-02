package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	if pos < 0 {
		return nil
	}
	for l != nil && pos > 0 {
		l = l.Next
		pos--
	}
	return l
}
