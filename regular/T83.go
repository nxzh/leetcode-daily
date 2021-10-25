package regular

import "fmt"

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var cach [201]bool
	cach[head.Val+100] = true
	for ptr, prev := head.Next, head; ptr != nil; ptr = ptr.Next {
		if cach[ptr.Val+100] {
			prev.Next = ptr.Next
		} else {
			cach[ptr.Val+100] = true
			prev = prev.Next
		}
	}
	return head
}

func T83() {

	res := deleteDuplicates(&ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, nil}}}}})
	fmt.Println(res)
}
