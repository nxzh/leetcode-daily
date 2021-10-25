package regular

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTarget(root *TreeNode, k int) bool {
	// iterate the tree in in-order order, for each element, find the target from tree
	return solution653(root, root, k)
}
func solution653(cur, root *TreeNode, k int) bool {
	if cur == nil {
		return false
	}
	return findTargetLeft(cur, root, k-cur.Val) || solution653(cur.Left, root, k) || solution653(cur.Right, root, k)
}
func findTargetLeft(cur, root *TreeNode, target int) bool {
	if root == nil {
		return false
	}
	if root.Val == target && root != cur {
		return true
	}
	if root.Val > target {
		return findTargetLeft(cur, root.Left, target)
	}
	if root.Val < target {
		return findTargetLeft(cur, root.Right, target)
	}
	return false
}

func T653() {
	//tree := &TreeNode{5,
	//	&TreeNode{3,
	//		&TreeNode{2, nil, nil},
	//		&TreeNode{4, nil, nil}},
	//
	//	&TreeNode{6, nil, &TreeNode{7, nil, nil}}}
	tree := &TreeNode{2, &TreeNode{0, &TreeNode{-4, nil, nil}, &TreeNode{1, nil, nil}}, &TreeNode{3, nil, nil}}
	res := findTarget(tree, -1)
	fmt.Println(res)
}
