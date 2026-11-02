package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	dummy := &NodeI{}
	tail := dummy
	for n1 != nil && n2 != nil {
		if n1.Data <= n2.Data {
			tail.Next = n1
			n1 = n1.Next
		} else {
			tail.Next = n2
			n2 = n2.Next
		}
		tail = tail.Next
	}
	if n1 != nil {
		tail.Next = n1
	} else {
		tail.Next = n2
	}
	return dummy.Next
}
