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
	for list != nil {
		fmt.Print(list.Val)
		fmt.Print("->")
		list = list.Next
	}
	fmt.Print("NULL")
}
