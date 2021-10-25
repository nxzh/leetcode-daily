package y2021m08

import (
	"math"
)

const MOD1000000007 = 1000000007

func maxProduct(root *TreeNode) int {
	var cache0819 []int64
	sum := traverseTree(root, &cache0819)
	var minDiff int64 = math.MaxInt64
	var minArr int64 = 0
	for _, val := range cache0819 {
		diff := int64(math.Abs(float64(sum>>1 - val)))
		if minDiff > diff {
			minDiff = diff
			minArr = val
		}
	}
	return int(minArr * (sum - minArr) % MOD1000000007)
}

func traverseTree(node *TreeNode, cache0819 *[]int64) int64 {
	if node == nil {
		return 0
	}
	sum := int64(node.Val)
	if node.Left != nil {
		resL := traverseTree(node.Left, cache0819)
		*cache0819 = append(*cache0819, resL)
		sum += resL
	}
	if node.Right != nil {
		resR := traverseTree(node.Right, cache0819)
		*cache0819 = append(*cache0819, resR)
		sum += resR
	}
	return sum
}
