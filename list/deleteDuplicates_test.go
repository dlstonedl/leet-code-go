package list

import (
	"fmt"
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	head := &ListNode{1, nil}
	head.Next = &ListNode{1, nil}
	head.Next.Next = &ListNode{2, nil}
	head.Next.Next.Next = &ListNode{3, nil}
	head.Next.Next.Next.Next = &ListNode{3, nil}

	node := deleteDuplicates(head)
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}
