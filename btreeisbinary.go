package piscine

func BTreeIsBinary(root *TreeNode) bool {
	return btreeIsBinaryBetween(root, nil, nil)
}

func btreeIsBinaryBetween(root *TreeNode, lower, upper *string) bool {
	if root == nil {
		return true
	}
	if lower != nil && root.Data < *lower || upper != nil && root.Data >= *upper {
		return false
	}
	return btreeIsBinaryBetween(root.Left, lower, &root.Data) &&
		btreeIsBinaryBetween(root.Right, &root.Data, upper)
}
