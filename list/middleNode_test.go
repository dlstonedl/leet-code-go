package list

import (
	"fmt"
	"testing"
)

func TestMiddleNode(t *testing.T) {
	head := &ListNode{1, nil}
	index := head
	for i := 2; i <= 5; i++ {
		next := &ListNode{i, nil}
		index.Next = next
		index = next
	}

	node := middleNode(head)
	if node.Val != 3 {
		fmt.Printf("val = %d, want 3", node.Val)
	}
}
