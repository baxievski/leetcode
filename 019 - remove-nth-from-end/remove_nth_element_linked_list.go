package removefromlinkedlist

// ListNode represents a node in a linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	return RemoveNthFromEnd(head, n)
}

// RemoveNthFromEnd removes the Nth element from the end of a linked list
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	current := head
	for i := 0; i < n; i++ {
		current = current.Next
	}
	if current == nil {
		return head.Next
	}
	nthFromEnd := head
	for current.Next != nil {
		current = current.Next
		nthFromEnd = nthFromEnd.Next
	}
	nthFromEnd.Next = nthFromEnd.Next.Next
	return head
}
