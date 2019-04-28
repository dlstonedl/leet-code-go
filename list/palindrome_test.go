package list

import "testing"

func TestIsPalindrome_1_2_2_2_1(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 1,
					},
				},
			},
		},
	}

	palindrome := isPalindrome(head)

	if palindrome != true {
		t.Errorf("palindrome is %v, expect %v", palindrome, true)
	}

}

func TestIsPalindrome_1_2_2_1(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
				},
			},
		},
	}

	palindrome := isPalindrome(head)

	if palindrome != true {
		t.Errorf("palindrome is %v, expect %v", palindrome, true)
	}
}

func TestIsPalindrome_1_2_1(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 1,
			},
		},
	}

	palindrome := isPalindrome(head)

	if palindrome != true {
		t.Errorf("palindrome is %v, expect %v", palindrome, true)
	}
}

func TestIsPalindrome_1_2(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}

	palindrome := isPalindrome(head)

	if palindrome != false {
		t.Errorf("palindrome is %v, expect %v", palindrome, false)
	}
}

func TestIsPalindrome_1(t *testing.T) {
	head := &ListNode{
		Val: 1,
	}

	palindrome := isPalindrome(head)

	if palindrome != true {
		t.Errorf("palindrome is %v, expect %v", palindrome, true)
	}
}
