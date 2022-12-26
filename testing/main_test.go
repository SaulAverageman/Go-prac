package main

import (
	"testing"
)

var tests = []struct {
	testname string
	dividend float32
	divisor  float32
	xError   bool
}{
	{"Normal-test", 5, 2, false},
	{"Divide by zero", 5, 0, true},
	{"Ans-check", 10.0, -3.0, false},
}

func TestDivide(t *testing.T) {
	for _, test := range tests {
		_, err := divide(test.dividend, test.divisor)

		if !test.xError {
			if err != nil {
				t.Error("Got unexpected error")
			}
		} else {
			if err == nil {
				t.Error("Should have got error")
			}
		}
	}
}
