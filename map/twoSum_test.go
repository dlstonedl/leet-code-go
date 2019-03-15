package mappractice

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	indexArray := twoSum([]int{2, 11, 7, 15}, 9)

	if indexArray[0] != 0 {
		t.Errorf("indexArray[0] = %d; want 0", indexArray[0])
	}

	if indexArray[1] != 2 {
		t.Errorf("indexArray[1] = %d; want 2", indexArray[1])
	}
}
