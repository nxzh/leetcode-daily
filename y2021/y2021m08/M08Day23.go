package y2021m08

import "fmt"

func findTarget(root *TreeNode, k int) bool {
	cache := make(map[int]bool)
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[i]
			val := node.Val
			if tryFindOtherwiseInsert(cache, val, k-val) {
				return true
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[size:]
	}
	return false
}

func tryFindOtherwiseInsert(cache map[int]bool, n int, target int) bool {
	if _, prs := cache[target]; prs {
		return true
	} else {
		cache[n] = true
	}
	return false
}

func Call0823() {
	tree := &TreeNode{5,
		&TreeNode{3,
			&TreeNode{2, nil, nil},
			&TreeNode{4, nil, nil}},
		&TreeNode{6, nil, &TreeNode{7, nil, nil}}}
	fmt.Println(findTarget(tree, 9))
}
