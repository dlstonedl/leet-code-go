package list

func hasCycle(head *ListNode) bool {
	low := head
	fast := head

	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		low = low.Next
		fast = fast.Next.Next
		if low.Val == fast.Val {
			return true
		}
	}

	return false
}
