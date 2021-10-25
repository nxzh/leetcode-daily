package regular

func getIntersectionNode(headA, headB *ListNode) *ListNode {

	len1 := lenOfList(headA)
	len2 := lenOfList(headB)
	ptr1 := headA
	ptr2 := headB
	if len1 > len2 {
		ptr1 = moveForward(ptr1, len1-len2)
	} else {
		ptr2 = moveForward(ptr2, len2-len1)
	}
	for ptr1 != nil && ptr2 != nil {
		if ptr1 == ptr2 {
			return ptr1
		}
		ptr1 = ptr1.Next
		ptr2 = ptr2.Next
	}
	return nil
}

func lenOfList(head *ListNode) int {
	ptr := head
	length := 0
	for ptr != nil {
		ptr = ptr.Next
		length++
	}
	return length
}
func moveForward(ptr *ListNode, n int) *ListNode {
	if n == 0 {
		return ptr
	}
	for n > 0 {
		ptr = ptr.Next
		n--
	}
	return ptr
}
