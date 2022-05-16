package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return nil
	}
	if string(root.Data) == elem {
		return root
	} else if string(root.Data) > elem {
		nd := root.Left
		if nd != nil {
			return BTreeSearchItem(nd, elem)
		} else {
			return nil
		}
	} else if string(root.Data) < elem {
		nd := root.Right
		if nd != nil {
			return BTreeSearchItem(nd, elem)
		} else {
			return nil
		}
	}
	return nil
}