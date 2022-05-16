package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	if root == nil {
		return rplc
	}
	if node == root {
		return rplc
	}
	nPrt := node.Parent
	if nPrt.Left.Data == node.Data {
		nPrt.Left = nil
	} else {
		nPrt.Right = nil
	}
	if rplc != nil {
		BTreeInsertData(root, rplc.Data)
	}
	return root
}