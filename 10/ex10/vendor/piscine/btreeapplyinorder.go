package piscine

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil || f == nil {
		return
	}
	l := root.Left
	if l != nil {
		BTreeApplyInorder(l, f)
	}
	f(root.Data)
	r := root.Right
	if r != nil {
		BTreeApplyInorder(r, f)
	}
}