package piscine

// no nil guard
func BTreeMax(root *TreeNode) *TreeNode {
	for root.Right != nil {
		root = root.Right
	}
	return root
}