package testrepo

import "testing"

func TestMissingValues(t *testing.T) {
	obj := GoTest{Val1: "A", Val2: "B"}
	if !obj.Test() {
		t.Error("missing values check failed")
	}
}
