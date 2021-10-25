package y2021m08

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const (
	LEFT = iota
	RIGHT
)

func pathSum(root *TreeNode, targetSum int) [][]int {
	var res [][]int
	solution0804(root, targetSum, 0, []int{}, &res)
	return res
}

func solution0804(node *TreeNode, targetSum int, currentSum int, path []int, res *[][]int) {
	if node == nil {
		return
	} else {
		path = append(path, node.Val)
		currentSum += node.Val
		if node.Left == nil && node.Right == nil {
			if currentSum == targetSum {
				*res = append(*res, append([]int{}, path...))
			}
		}
		if node.Left != nil {
			solution0804(node.Left, targetSum, currentSum, path, res)
		}
		if node.Right != nil {
			solution0804(node.Right, targetSum, currentSum, path, res)
		}
	}
}
