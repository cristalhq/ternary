package ternary_test

import (
	"testing"

	. "github.com/cristalhq/ternary"
)

func TestAndBochvar(t *testing.T) {
	tcs := []struct {
		a, b, x Value
	}{
		{False, False, False},
		{False, Unknown, Unknown},
		{False, True, False},

		{Unknown, False, Unknown},
		{Unknown, Unknown, Unknown},
		{Unknown, True, Unknown},

		{True, False, False},
		{True, Unknown, Unknown},
		{True, True, True},
	}

	for _, tc := range tcs {
		res := AndB(tc.a, tc.b)
		if res != tc.x {
			t.Errorf("(%v,%v) want %v got %v", tc.a, tc.b, tc.x, res)
		}
	}
}

func TestOrBochvar(t *testing.T) {
	tcs := []struct {
		a, b, x Value
	}{
		{False, False, False},
		{False, Unknown, Unknown},
		{False, True, True},

		{Unknown, False, Unknown},
		{Unknown, Unknown, Unknown},
		{Unknown, True, Unknown},

		{True, False, True},
		{True, Unknown, Unknown},
		{True, True, True},
	}

	for _, tc := range tcs {
		res := OrB(tc.a, tc.b)
		if res != tc.x {
			t.Errorf("(%v,%v) want %v got %v", tc.a, tc.b, tc.x, res)
		}
	}
}

func TestImpBochvar(t *testing.T) {
	tcs := []struct {
		a, b, x Value
	}{
		{False, False, True},
		{False, Unknown, Unknown},
		{False, True, True},

		{Unknown, False, Unknown},
		{Unknown, Unknown, Unknown},
		{Unknown, True, True},

		{True, False, False},
		{True, Unknown, Unknown},
		{True, True, True},
	}

	for _, tc := range tcs {
		res := ImpB(tc.a, tc.b)
		if res != tc.x {
			t.Errorf("(%v,%v) want %v got %v", tc.a, tc.b, tc.x, res)
		}
	}
}
