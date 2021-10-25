package y2021m09

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func copyTree(from *TreeNode, to *TreeNode) {
	if from == nil {
		return
	}
	if from.Left != nil {
		to.Left = &TreeNode{Val: from.Left.Val}
		copyTree(from.Left, to.Left)
	}
	if from.Right != nil {
		to.Right = &TreeNode{Val: from.Right.Val}
		copyTree(from.Right, to.Right)
	}
}

func duplicateTree(node *TreeNode) *TreeNode {
	copied := &TreeNode{Val: node.Val}
	copyTree(node, copied)
	return copied
}

func generateTrees(n int) []*TreeNode {
	root := &TreeNode{Val: 1}
	track := []*TreeNode{root}
	for i := 2; i <= n; i++ {
		lenTrack := len(track)
		for j := 0; j < lenTrack; j++ {
			node := track[j]
			dup := duplicateTree(node)
			track = append(track, &TreeNode{Val: i, Left: dup})
			generateTreesInternal(i, track[j], node.Val, track[j], &track)
		}
		track = track[lenTrack:]
	}
	return track
}

func generateTreesInternal(n int, node *TreeNode, parent int, root *TreeNode, track *[]*TreeNode) {
	dupTree := duplicateTree(root)
	newNode := &TreeNode{Val: n}
	ptr := dupTree
	for {
		if ptr.Val == parent {
			break
		}
		ptr = ptr.Right
	}
	oldRight := ptr.Right
	ptr.Right = newNode
	newNode.Left = oldRight
	*track = append(*track, dupTree)
	if node.Right != nil {
		generateTreesInternal(n, node.Right, node.Right.Val, root, track)
	}
}

func Call02() {
	//res := generateTrees(1)
	res := generateTrees(4)
	fmt.Println(res)
}
