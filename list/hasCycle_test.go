package list

import (
	"fmt"
	"testing"
)

func TestHasCycle(t *testing.T) {
	head := &ListNode{Val: 3}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 0}
	head.Next.Next.Next = &ListNode{Val: -4}
	head.Next.Next.Next.Next = head.Next

	if cycle := hasCycle(head); !cycle {
		fmt.Printf("cycle = %v", cycle)
	}
}
