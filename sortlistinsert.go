package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	node := &NodeI{Data: data_ref}
	if l == nil || data_ref < l.Data {
		node.Next = l
		return node
	}
	position := l
	for position.Next != nil && position.Next.Data <= data_ref {
		position = position.Next
	}
	node.Next = position.Next
	position.Next = node
	return l
}
