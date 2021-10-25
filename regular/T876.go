package regular

import "fmt"

func middleNode(head *ListNode) *ListNode {
	// 1. slow, fast := head, head
	// 2. slow+1 every time, fast+2 every time
	// 3. when fast .next || fast.next.next  is null, slow is pointing to the middle
	slow, fast := head, head
	for {
		if fast.Next == nil {
			return slow
		} else if fast.Next.Next == nil {
			return slow.Next
		}
		fast = fast.Next.Next
		slow = slow.Next
	}

}

func T876() {
	list := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	//list := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}}
	res := middleNode(list)
	fmt.Println(res)
}
