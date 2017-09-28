package main

import (
	"playground/kata/pancake2/waiters"
	"testing"
)

var testData = map[string]struct {
	stack string
	flips int
}{
	"NoFlips":     {"+", 0},
	"OneFlip":     {"-", 1},
	"PosNeg":      {"+-", 2},
	"Alternating": {"-+-+-+-+-", 9},
	"Final":       {"-++++++++++-", 3},
}

func Test(t *testing.T) {
	for key, values := range testData {
		waiter := waiters.New()
		flips := waiter.ArrangeStack(values.stack)
		if flips != values.flips {
			t.Errorf("[%s] expected %d flips, got %d", key, values.flips, flips)
		}
	}
}
