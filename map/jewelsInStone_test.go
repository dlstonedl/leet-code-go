package mappractice

import "testing"

func TestNumJewelsInStones(t *testing.T) {
	result := numJewelsInStones("aA", "aAAbbbb")

	if result != 3 {
		t.Errorf("result = %d; want 3", result)
	}
}

func TestNumJewelsInStonesThoughStringsIndex(t *testing.T) {
	result := numJewelsInStonesThoughStringsIndex("aA", "aAAbbbb")

	if result != 3 {
		t.Errorf("result = %d; want 3", result)
	}
}
