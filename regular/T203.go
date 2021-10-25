package regular

func removeElements(head *ListNode, val int) *ListNode {
	// 1. dummy->Next = head
	// 2. if ptr.Next.Val == val, ptr = ptr.Next.Next
	dummy := &ListNode{}
	dummy.Next = head
	ptr := dummy
	for ptr != nil && ptr.Next != nil {
		if ptr.Next.Val == val {
			ptr.Next = ptr.Next.Next
		} else {
			ptr = ptr.Next
		}
	}
	return dummy.Next
}
