package list

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	pre := head
	next := head.Next
	for next != nil {
		if pre.Val == next.Val {
			pre.Next = next.Next
			next = next.Next
		} else {
			pre = next
			next = next.Next
		}
	}

	return head
}
