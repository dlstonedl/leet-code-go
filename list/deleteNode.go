package list

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

func deleteNode1(head, node *ListNode) *ListNode {
	index := head
	temp := head
	for index.Next != nil {
		if index.Val == node.Val {
			temp.Next = index.Next
		}
		temp = index
		index = index.Next
	}
	return head
}
