package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	for l.Head != nil && l.Head.Data == data_ref {
		l.Head = l.Head.Next
	}
	if l.Head == nil {
		l.Tail = nil
		return
	}

	for node := l.Head; node.Next != nil; {
		if node.Next.Data == data_ref {
			node.Next = node.Next.Next
		} else {
			node = node.Next
		}
	}
	l.Tail = l.Head
	for l.Tail.Next != nil {
		l.Tail = l.Tail.Next
	}
}
