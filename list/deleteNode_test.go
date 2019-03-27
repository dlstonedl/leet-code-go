package list

import (
	"fmt"
	"testing"
)

func TestDeleteNode(t *testing.T) {
	head := &ListNode{4, nil}
	index := head
	index.Next = &ListNode{5, nil}
	index = index.Next
	index.Next = &ListNode{1, nil}
	index = index.Next
	index.Next = &ListNode{9, nil}

	head = deleteNode1(head, &ListNode{5, nil})
	for head.Next != nil {
		if head.Val == 5 {
			fmt.Printf("error")
		}
		head = head.Next
	}
}
