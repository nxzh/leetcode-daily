package regular

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	ptr1 := l1
	ptr2 := l2
	head := &ListNode{}
	ptr := head
	for {
		if ptr1 == nil && ptr2 == nil {
			break
		}
		val1 := 0
		if ptr1 != nil {
			val1 = ptr1.Val
			ptr1 = ptr1.Next
		}
		val2 := 0
		if ptr2 != nil {
			val2 = ptr2.Val
			ptr2 = ptr2.Next
		}
		val := val1 + val2 + carry
		if val >= 10 {
			carry = 1
			val -= 10
		} else {
			carry = 0
		}
		ptr.Next = &ListNode{Val: val}
		ptr = ptr.Next
	}
	if carry > 0 {
		ptr.Next = &ListNode{Val: carry}
	}
	return head.Next
}
