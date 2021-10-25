package regular

func isPalindrome234(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	var previous *ListNode = nil
	current := slow
	next := current
	for current != nil {
		next = current.Next
		current.Next = previous
		previous = current
		current = next
	}
	ptrNew := previous
	ptr := head
	for ptrNew != nil && ptr != nil {
		if ptr.Val != ptrNew.Val {
			return false
		}
		ptr = ptr.Next
		ptrNew = ptrNew.Next
	}
	return true
}
