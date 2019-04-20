package ternary_test

import (
	"testing"

	. "github.com/cristalhq/ternary"
)

func TestNeg(t *testing.T) {
	tcs := []struct {
		a Value
		x Value
	}{
		{False, True},
		{Unknown, Unknown},
		{True, False},
	}

	for _, tc := range tcs {
		res := Not(tc.a)
		if res != tc.x {
			t.Errorf("(%v) want %v got %v", tc.a, tc.x, res)
		}
	}
}

func TestAnd(t *testing.T) {
	tcs := []struct {
		a Value
		b Value
		x Value
	}{
		{False, False, False},
		{False, Unknown, False},
		{False, True, False},

		{Unknown, False, False},
		{Unknown, Unknown, Unknown},
		{Unknown, True, Unknown},

		{True, False, False},
		{True, Unknown, Unknown},
		{True, True, True},
	}

	for _, tc := range tcs {
		res := And(tc.a, tc.b)
		if res != tc.x {
			t.Errorf("(%v,%v) want %v got %v", tc.a, tc.b, tc.x, res)
		}
	}
}

func TestOr(t *testing.T) {
	tcs := []struct {
		a Value
		b Value
		x Value
	}{
		{False, False, False},
		{False, Unknown, Unknown},
		{False, True, True},

		{Unknown, False, Unknown},
		{Unknown, Unknown, Unknown},
		{Unknown, True, True},

		{True, False, True},
		{True, Unknown, True},
		{True, True, True},
	}

	for _, tc := range tcs {
		res := Or(tc.a, tc.b)
		if res != tc.x {
			t.Errorf("(%v,%v) want %v got %v", tc.a, tc.b, tc.x, res)
		}
	}
}

func TestImp(t *testing.T) {
	tcs := []struct {
		a Value
		b Value
		x Value
	}{
		{False, False, True},
		{False, Unknown, True},
		{False, True, True},

		{Unknown, False, Unknown},
		{Unknown, Unknown, Unknown},
		{Unknown, True, True},

		{True, False, False},
		{True, Unknown, Unknown},
		{True, True, True},
	}

	for _, tc := range tcs {
		res := Imp(tc.a, tc.b)
		if res != tc.x {
			t.Errorf("(%v,%v) want %v got %v", tc.a, tc.b, tc.x, res)
		}
	}
}

func TestImpL(t *testing.T) {
	tcs := []struct {
		a Value
		b Value
		x Value
	}{
		{False, False, True},
		{False, Unknown, True},
		{False, True, True},

		{Unknown, False, Unknown},
		{Unknown, Unknown, True},
		{Unknown, True, True},

		{True, False, False},
		{True, Unknown, Unknown},
		{True, True, True},
	}

	for _, tc := range tcs {
		res := ImpL(tc.a, tc.b)
		if res != tc.x {
			t.Errorf("(%v,%v) want %v got %v", tc.a, tc.b, tc.x, res)
		}
	}
}

func TestEqv(t *testing.T) {
	tcs := []struct {
		a Value
		b Value
		x Value
	}{
		{False, False, True},
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
		res := Eqv(tc.a, tc.b)
		if res != tc.x {
			t.Errorf("(%v,%v) want %v got %v", tc.a, tc.b, tc.x, res)
		}
	}
}
