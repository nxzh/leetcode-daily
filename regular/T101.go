package regular

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}
	return solution101(root.Left, root.Right)
}

func solution101(l *TreeNode, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	} else if l != nil && r != nil && l.Val == r.Val {
		return solution101(l.Left, r.Right) && solution101(l.Right, r.Left)
	} else {
		return false
	}
}
