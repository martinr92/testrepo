package testrepo

import "testing"

func TestMissingValues(t *testing.T) {
	obj := GoTest{Val1: "A", Val2: "B"}
	if !obj.Test() {
		t.Error("missing values check failed")
	}
}

func TestMissingValueFail1(t *testing.T) {
	obj := GoTest{Val1: "A"}
	if obj.Test() {
		t.Error("val2 is missing! function Test should return false")
	}
}

func TestMissingValueFail2(t *testing.T) {
	obj := GoTest{Val2: "B"}
	if obj.Test() {
		t.Error("val1 is missing! function Test should return false")
	}
}
