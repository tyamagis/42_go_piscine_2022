ackage piscine

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	var new TreeNode
	new.Data = data
	if root == nil {
		root = &new
		return &new
	}
	node := root
	for node != nil {
		if node.Data > new.Data {
			if node.Left == nil {
				node.Left = &new
				new.Parent = node
				return &new
			} else {
				node = node.Left
			}
		} else if node.Data < new.Data {
			if node.Right == nil {
				node.Right = &new
				new.Parent = node
				return &new
			} else {
				node = node.Right
			}
		}
	}
	node = &new
	return &new
}
