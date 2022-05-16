package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lv := 1
	if root.Left == nil && root.Right == nil {
		return lv
	}
	maxLv := lv
	nd := root.Left
	if nd != nil {
		lv += BTreeLevelCount(nd)
		if maxLv < lv {
			maxLv = lv
		}
		lv = 1
	}
	nd = root.Right
	if nd != nil {
		lv += BTreeLevelCount(nd)
		if maxLv < lv {
			maxLv = lv
		}
		lv = 1
	}
	return maxLv
}