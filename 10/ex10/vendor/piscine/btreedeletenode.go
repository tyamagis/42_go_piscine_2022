package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil || node == nil {
		return root
	}
	if node.Left == nil && node.Right == nil { // no child
		nPrt := node.Parent
		if nPrt == nil { // no parent, no child
			return nil
		}
		if nPrt.Left != nil && nPrt.Left.Data == node.Data {
			nPrt.Left = nil
		}
		if nPrt.Right != nil && nPrt.Right.Data == node.Data {
			nPrt.Right = nil
		}
		return root
	} else if node.Left == nil || node.Right == nil { // has a child
		nPrt := node.Parent
		nCld := node.Left
		if nCld == nil {
			nCld = node.Right
		}
		nCld.Parent = nPrt
		if nPrt != nil {
			if nPrt.Left != nil && nPrt.Left.Data == node.Data {
				nPrt.Left = nCld
			} else {
				nPrt.Right = nCld
			}
		}
		return root
	} else { // has two child
		
	}
}