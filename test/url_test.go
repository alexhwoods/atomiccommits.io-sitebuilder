package test

import "testing"

func Canary(t *testing.T) {
	x := 4

	if x == 4 {
		t.Errorf("x should not equal 4, it is")
	}
}
