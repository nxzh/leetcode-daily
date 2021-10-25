package regular

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	var queue []int
	for head != nil {
		queue = append(queue, head.Val)
		head = head.Next
	}
	return solution109(queue)
}

func solution109(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}
	return toTree109(nums)
}

func toTree109(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	var ret TreeNode
	ret.Val = nums[len(nums)/2]
	ret.Left = toTree109(nums[0 : len(nums)/2])
	ret.Right = toTree109(nums[len(nums)/2+1:])
	return &ret
}
