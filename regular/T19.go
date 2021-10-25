package regular

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 1. slow, fast = head, head+n
	// 2. when fast.Next == nil, slow is n-1, then slow = slow.next.next
	dummy := &ListNode{0, head}
	count, slow, fast := 0, dummy, dummy
	for count < n {
		fast = fast.Next
		count++
	}
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}

func T19() {
	list := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	res := removeNthFromEnd(list, 1)
	fmt.Println(res)
}
