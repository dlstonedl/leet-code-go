package list

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	first := head
	head = head.Next
	first.Next = nil

	for head != nil {
		temp := head.Next
		head.Next = first
		first = head
		head = temp
	}
	return first
}
