package regular

func isValidBST(root *TreeNode) bool {
	// solution 1
	//return isValid98(root, math.MinInt64, math.MaxInt64)

	// solution 2
	result := []int{}
	solution98(root, &result)
	for i := 0; i < len(result)-1; i++ {
		if result[i+1] <= result[i] {
			return false
		}
	}
	return true
}
func isValid98(node *TreeNode, lower int, upper int) bool {
	if node == nil {
		return true
	}
	if node.Val <= lower || node.Val >= upper {
		return false
	}

	return isValid98(node.Left, lower, node.Val) && isValid98(node.Right, node.Val, upper)
}

func solution98(node *TreeNode, ary *[]int) {
	if node.Left != nil {
		solution98(node.Left, ary)
	}
	*ary = append(*ary, node.Val)
	if node.Right != nil {
		solution98(node.Right, ary)
	}
}
