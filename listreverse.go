package piscine

func ListReverse(l *List) {
	l.Tail = l.Head
	var previous *NodeL
	for node := l.Head; node != nil; {
		next := node.Next
		node.Next = previous
		previous = node
		node = next
	}
	l.Head = previous
}
