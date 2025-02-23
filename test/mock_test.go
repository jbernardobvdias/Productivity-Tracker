package test

import "testing"

func TestMock(t *testing.T) {
	a := 1

	if a != 1 {
		t.Fatalf("")
	}
}
