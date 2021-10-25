package regular

func postorderTraversal(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}
	solution145(root, &result)
	return result
}

func solution145(node *TreeNode, result *[]int) {
	if node.Left != nil {
		solution145(node.Left, result)
	}
	if node.Right != nil {
		solution145(node.Right, result)
	}
	*result = append(*result, node.Val)
}
