package regular

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	merged := new(ListNode)
	ptr := merged
	for l1 != nil && l2 != nil {
		node := new(ListNode)
		if l1.Val < l2.Val {
			node.Val = l1.Val
			l1 = l1.Next
		} else {
			node.Val = l2.Val
			l2 = l2.Next
		}
		ptr.Next = node
		ptr = ptr.Next
	}
	if l1 != nil {
		mergeRemain(l1, ptr)
	}
	if l2 != nil {
		mergeRemain(l2, ptr)
	}
	return merged.Next
}

func mergeRemain(l *ListNode, merged *ListNode) {
	for l != nil {
		node := new(ListNode)
		node.Val = l.Val
		l = l.Next
		merged.Next = node
		merged = merged.Next
	}
}
