package testrepo

import (
	"fmt"
)

// GoTest is a test structure.
type GoTest struct {
	Val1 string
	Val2 string
}

// Test is a function, that does absolutely nothing.
func (gt *GoTest) Test() bool {
	if gt.Val1 == "" {
		fmt.Println("val 1 is empty")
		return false
	}

	if gt.Val2 == "" {
		fmt.Println("val 2 is empty")
		return false
	}

	return true
}
