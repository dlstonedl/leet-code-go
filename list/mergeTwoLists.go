package list

func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	head := &ListNode{}
	index := head
	for {
		if l1 == nil || l2 == nil {
			break
		}

		if l1.Val <= l2.Val {
			index.Next = l1
			index = index.Next
			l1 = l1.Next
		} else {
			index.Next = l2
			index = index.Next
			l2 = l2.Next
		}
	}

	if l1 != nil {
		index.Next = l1
	}

	if l2 != nil {
		index.Next = l2
	}

	return head.Next
}
