package piscine

func BTreeApplyPreorder(root *TreeNode, f func(...interface{}) (int, error)){
	if root == nil || f == nil {
		return
	}
	f(root.Data)
	nd := root.Left
	if nd != nil {
		BTreeApplyPreorder(nd, f)
	}
	nd = root.Right
	if nd != nil {
		BTreeApplyPreorder(nd, f)
	}
}