package regular

func insertIntoBST(node *TreeNode, val int) *TreeNode {
	if node == nil {
		return &TreeNode{Val: val}
	}
	if node.Val > val {
		node.Left = insertIntoBST(node.Left, val)
	}
	if node.Val < val {
		node.Right = insertIntoBST(node.Right, val)
	}
	return node
}
