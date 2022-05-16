package piscine

func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return false
	}
	if root.Left != nil {
		if root.Data < root.Left.Data || BTreeIsBinary(root.Left) == false {
			return false
		}
	}
	if root.Right != nil {
		if root.Data > root.Right.Data || BTreeIsBinary(root.Right) == false {
			return false
		}
	}
	return true
}