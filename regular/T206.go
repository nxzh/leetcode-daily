package regular

func reverseList(head *ListNode) *ListNode {

	current := head
	var previous, next *ListNode

	// for current is not null
	// next = current.next
	// current.next = previous
	// previous = current
	// current = current.next

	for current != nil {
		next = current.Next
		current.Next = previous
		previous = current
		current = next
	}
	return previous
}
