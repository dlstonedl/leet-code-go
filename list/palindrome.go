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
