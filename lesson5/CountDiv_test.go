package lesson5

import (
	"testing"
)

func TestSolution(t *testing.T) {
	result := Solution(6, 11, 2)
	if result != 3 {
		t.Errorf("expect 3, but found %s", result)
	}

	result = Solution(11, 345, 17)
	if result != 20 {
		t.Errorf("expect 20, but found %s", result)
	}

	result = Solution(0, 0, 11)
	if result != 1 {
		t.Errorf("expect 1, but found %s", result)
	}

	result = Solution(10, 10, 5)
	if result != 1 {
		t.Errorf("expect 1, but found %s", result)
	}

	result = Solution(11, 14, 2)
	if result != 2 {
		t.Errorf("expect 2, but found %s", result)
	}
}
