package list

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	if head.Next == nil {
		return true
	}

	fast, low := head, head
	var values []int
	values = append(values, low.Val)

	for fast.Next != nil && fast.Next.Next != nil {
		low = low.Next
		values = append(values, low.Val)
		fast = fast.Next.Next
	}

	if fast.Next != nil {
		low = low.Next
	}

	for low != nil {
		if low.Val != values[len(values)-1] {
			return false
		} else {
			values = values[0 : len(values)-1]
		}
		low = low.Next
	}

	if len(values) != 0 {
		return false
	}

	return true
}

func isPalindrome2(head *ListNode) bool {
	if head == nil {
		return true
	}

	if head.Next == nil {
		return true
	}

	if head.Next.Next == nil &&
		head.Val == head.Next.Val {
		return true
	}

	low, fast := head, head
	index := low.Next

	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		low = index
		index = index.Next

		low.Next = head
		if head.Next == low {
			head.Next = nil
		}
		head = low
	}

	if fast.Next == nil {
		head = head.Next
	}

	for head != nil {

		if head.Val != index.Val {
			return false
		}

		head = head.Next
		index = index.Next
	}

	return true
}
