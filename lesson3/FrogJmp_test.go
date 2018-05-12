package lesson3

import (
	"testing"
)

func TestSolution2(t *testing.T) {
	result := Solution2(10, 85, 30)
	if result != 3{
		t.Errorf("expect 3, but found %v", result)
	}
}