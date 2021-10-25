package regular

func preorderTraversal(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}
	solution144(root, &result)
	return result
}

func solution144(node *TreeNode, result *[]int) {
	*result = append(*result, node.Val)
	if node.Left != nil {
		solution144(node.Left, result)
	}
	if node.Right != nil {
		solution144(node.Right, result)
	}
}
