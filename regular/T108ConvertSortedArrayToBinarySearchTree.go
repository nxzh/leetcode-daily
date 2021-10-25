package regular

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}
	return toTree(nums)
}

func toTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	var ret TreeNode
	ret.Val = nums[len(nums)/2]
	ret.Left = toTree(nums[0 : len(nums)/2])
	ret.Right = toTree(nums[len(nums)/2+1:])
	return &ret
}
