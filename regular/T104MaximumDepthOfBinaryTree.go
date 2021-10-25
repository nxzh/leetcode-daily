package regular

import "math"

//func maxDepth(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	var queue []*TreeNode
//	queue = append(queue, root)
//	levels := 0
//	for len(queue) > 0 {
//		size := len(queue)
//		levels++
//		for i := 0; i < size; i++ {
//			node := queue[0]
//			queue = queue[1:]
//			if node.Left != nil {
//				queue = append(queue, node.Left)
//			}
//			if node.Right != nil {
//				queue = append(queue, node.Right)
//			}
//		}
//	}
//	return levels
//}

func maxDepth(root *TreeNode) int {
	return solution104(root)
}
func solution104(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return int(math.Max(float64(1+solution104(node.Left)), float64(1+solution104(node.Right))))
}
