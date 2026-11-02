package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	var sorted *NodeI
	for l != nil {
		node := l
		l = l.Next
		if sorted == nil || node.Data < sorted.Data {
			node.Next = sorted
			sorted = node
			continue
		}
		position := sorted
		for position.Next != nil && position.Next.Data <= node.Data {
			position = position.Next
		}
		node.Next = position.Next
		position.Next = node
	}
	return sorted
}
