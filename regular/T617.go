package regular

import "fmt"

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	// iterate level by level
	if root1 == nil && root2 == nil {
		return nil
	}
	merged := &TreeNode{}
	queue := [][]*TreeNode{{root1, root2, merged}}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node1, node2, node := queue[i][0], queue[i][1], queue[i][2]
			var l1, l2, r1, r2 *TreeNode
			if node1 == nil {
				node.Val = node2.Val
				l1, r1, l2, r2 = nil, nil, node2.Left, node2.Right
			} else if node2 == nil {
				node.Val = node1.Val
				l1, r1, l2, r2 = node1.Left, node1.Right, nil, nil
			} else {
				node.Val = node1.Val + node2.Val
				l1, r1, l2, r2 = node1.Left, node1.Right, node2.Left, node2.Right
			}
			if l1 == nil {
				node.Left = l2
			} else if l2 == nil {
				node.Left = l1
			} else {
				node.Left = &TreeNode{}
				queue = append(queue, []*TreeNode{l1, l2, node.Left})
			}
			if r1 == nil {
				node.Right = r2
			} else if r2 == nil {
				node.Right = r1
			} else {
				node.Right = &TreeNode{}
				queue = append(queue, []*TreeNode{r1, r2, node.Right})
			}
		}
		queue = queue[size:]
	}
	return merged
}
func mergeTreesDfs(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	// from the root nodes of two trees to bottom
	// if two nodes are both nil, then return
	// 		else both of them are not nil, add two values of two nodes. if one node is nil, its value is 0
	// 		do the above logic for the left nodes of the two trees. if one node is nil, copy the left of the other tree to the merging-tree
	// 		do the above logic for the right node
	if root1 == nil && root2 == nil {
		return nil
	}
	merged := &TreeNode{}
	solution617(root1, root2, merged)
	return merged
}

func solution617(node1 *TreeNode, node2 *TreeNode, node *TreeNode) {
	var l1, l2, r1, r2 *TreeNode
	if node1 == nil {
		node.Val = node2.Val
		l1, r1, l2, r2 = nil, nil, node2.Left, node2.Right
	} else if node2 == nil {
		node.Val = node1.Val
		l1, r1, l2, r2 = node1.Left, node1.Right, nil, nil
	} else {
		node.Val = node1.Val + node2.Val
		l1, r1, l2, r2 = node1.Left, node1.Right, node2.Left, node2.Right
	}
	if l1 != nil || l2 != nil {
		node.Left = &TreeNode{}
		solution617(l1, l2, node.Left)
	}
	if r1 != nil || r2 != nil {
		node.Right = &TreeNode{}
		solution617(r1, r2, node.Right)
	}
}

func T617() {
	//root1 := &TreeNode{1,
	//	&TreeNode{3, &TreeNode{5, nil, nil}, nil},
	//	&TreeNode{2, nil, nil}}
	//root2 := &TreeNode{2,
	//	&TreeNode{1, nil, &TreeNode{4, nil, nil}},
	//	&TreeNode{3, nil, &TreeNode{7, nil, nil}}}
	root1 := &TreeNode{1,
		&TreeNode{2, &TreeNode{3, nil, nil}, nil}, nil}
	root2 := &TreeNode{1, nil,
		&TreeNode{2, nil, &TreeNode{3, nil, nil}}}
	res := mergeTrees(root1, root2)
	fmt.Println(res)
}
