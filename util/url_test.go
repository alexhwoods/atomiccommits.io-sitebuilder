package util

import "testing"

func TestCanary(t *testing.T) {
	x := 4

	if x == 4 {
		t.Errorf("x should not equal 4, it is")
	}
}
