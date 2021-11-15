package regular

import "fmt"

func reverseEvenLengthGroups(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	ptr := head
	length := 0
	for ptr != nil {
		length++
		ptr = ptr.Next
	}
	if length == 1 {
		return head
	}
	ptr = &ListNode{
		Val:  0,
		Next: head,
	}
	count := 0
	for i := 1; count < length; i++ {
		moveCount := i
		if count+i >= length {
			moveCount = length - count
		}
		if moveCount%2 == 0 {
			reverseLink(ptr, moveCount)
		}
		for j := 0; j < moveCount; j++ {
			ptr = ptr.Next
		}
		count += i
	}
	return head
}

func reverseLink(from *ListNode, count int) {
	tail := from
	tailCounter := count + 1
	for tailCounter > 0 {
		tail = tail.Next
		tailCounter--
	}
	var prev = tail
	var cur = from.Next
	var next *ListNode = nil
	reverseCount := count
	for reverseCount > 0 {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
		reverseCount--
	}
	from.Next = prev
}

func T5927() {
	list := &ListNode{Val: 0, Next: &ListNode{2, &ListNode{4, &ListNode{1, &ListNode{3, nil}}}}}
	//list := &ListNode{Val: 1, Next: &ListNode{2, &ListNode{3, nil}}}
	reverseEvenLengthGroups(list)
	fmt.Println(list)
}
