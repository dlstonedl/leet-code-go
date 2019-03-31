package list

import (
	"fmt"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	l1 := &ListNode{1, nil}
	l1.Next = &ListNode{2, nil}
	l1.Next.Next = &ListNode{4, nil}

	l2 := &ListNode{1, nil}
	l2.Next = &ListNode{3, nil}
	l2.Next.Next = &ListNode{4, nil}

	head := mergeTwoLists(l1, l2)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
