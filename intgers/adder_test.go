package intgers

import "testing"

func TestAdder(t *testing.T) {
	sum := add(2,2)
	expected:= 4

	if sum != expected {
		t.Errorf("sum %d expected %d",sum,expected)
	}

}
