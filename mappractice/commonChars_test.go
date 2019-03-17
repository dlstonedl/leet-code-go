package mappractice

import "testing"

func TestCommonChars(t *testing.T) {
	result := commonChars([]string{"bella", "label", "roller"})

	if len(result) != 3 {
		t.Errorf("result len = %d; want 3;", len(result))
	}
}

func TestCommonCharsThoughArray(t *testing.T) {
	result := commonCharsThoughArray([]string{"bella", "label", "roller"})

	if len(result) != 3 {
		t.Errorf("result len = %d; want 3;", len(result))
	}
}
