package regular

func hasPathSum(node *TreeNode, targetSum int) bool {
	if node == nil {
		return false
	}
	if node.Left == nil && node.Right == nil && node.Val == targetSum {
		return true
	}

	nextTarget := targetSum - node.Val
	if node.Left != nil && hasPathSum(node.Left, nextTarget) {
		return true
	}
	if node.Right != nil && hasPathSum(node.Right, nextTarget) {
		return true
	}
	return false
}
