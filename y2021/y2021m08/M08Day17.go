package y2021m08

func goodNodes(root *TreeNode) int {
	return goodNodesInterval(root, -10001)
}

func goodNodesInterval(node *TreeNode, pathMax int) int {
	if node == nil {
		return 0
	}
	count := 0
	if node.Val >= pathMax {
		count++
		pathMax = node.Val
	}
	return count + goodNodesInterval(node.Left, pathMax) + goodNodesInterval(node.Right, pathMax)
}
