package regular

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const (
	LEFT = iota
	RIGHT
)

func inorderTraversal(root *TreeNode) []int {
	parent := root
	var result []int
	queue := []*TreeNode{root}
	order := LEFT
	for len(queue) > 0 {
		if order == LEFT && parent.Left != nil {
			queue = append(queue, parent.Left)
			parent = parent.Left
			continue
		}
		result = append(result, parent.Val)
		if parent.Right != nil {
			queue = append(queue, parent.Right)
			parent = parent.Right
			continue
		}
		parent = queue[len(queue)-1]
		order = RIGHT
	}
	return result
}

func Contains(queue []*TreeNode, x *TreeNode) bool {
	for _, n := range queue {
		if x == n {
			return true
		}
	}
	return false
}

//func inorderTraversal(root *TreeNode) []int {
//	var result []int
//	if root == nil {
//		return result
//	}
//
//	solution94(root, &result)
//	return result
//}
//
//func solution94(node *TreeNode, result *[]int) {
//	if node.Left != nil {
//		solution94(node.Left, result)
//	}
//	*result = append(*result, node.Val)
//	if node.Right != nil {
//		solution94(node.Right, result)
//	}
//}
