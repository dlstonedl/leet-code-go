package mappractice

import "testing"

func TestRepeatedNTimes(t *testing.T) {
	value := repeatedNTimes([]int{5, 1, 5, 2, 5, 3, 5, 4})

	if value != 5 {
		t.Errorf("value = %d; want 5", value)
	}
}
