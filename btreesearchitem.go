package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	for root != nil {
		if elem == root.Data {
			return root
		}
		if elem < root.Data {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return nil
}
