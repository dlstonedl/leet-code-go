package list

import (
	"fmt"
	"testing"
)

func TestReverseList(t *testing.T) {
	head := &ListNode{1, nil}
	index := head
	for i := 2; i <= 5; i++ {
		next := &ListNode{i, nil}
		index.Next = next
		index = next
	}

	list := reverseList(head)
	if list.Val != 5 {
		fmt.Printf("list.val = %d, want 5", list.Val)
	}
}

func TestReverseListByRecursion(t *testing.T) {
	head := &ListNode{1, nil}
	index := head
	for i := 2; i <= 5; i++ {
		next := &ListNode{i, nil}
		index.Next = next
		index = next
	}

	list := reverseListByRecursion(head)
	if list.Val != 5 {
		fmt.Printf("list.val = %d, want 5", list.Val)
	}
}
