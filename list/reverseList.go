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

func reverseListByRecursion(head *ListNode) *ListNode {
	_, root := recursion(head)
	return root
}

func recursion(head *ListNode) (cur, root *ListNode) {
	if head == nil {
		return nil, nil
	}

	prev, root := recursion(head.Next)
	if prev == nil {
		return head, head
	}

	cur = &ListNode{Val: head.Val}
	prev.Next = cur

	return cur, root
}
