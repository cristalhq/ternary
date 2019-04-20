package ternary_test

import (
	"testing"

	. "github.com/cristalhq/ternary"
)

func TestMA(t *testing.T) {
	tcs := []struct {
		a, x Value
	}{
		{False, False},
		{Unknown, True},
		{True, True},
	}

	for _, tc := range tcs {
		res := MA(tc.a)
		if res != tc.x {
			t.Errorf("(%v) want %v got %v", tc.a, tc.x, res)
		}
	}
}

func TestLA(t *testing.T) {
	tcs := []struct {
		a, x Value
	}{
		{False, False},
		{Unknown, False},
		{True, True},
	}

	for _, tc := range tcs {
		res := LA(tc.a)
		if res != tc.x {
			t.Errorf("(%v) want %v got %v", tc.a, tc.x, res)
		}
	}
}

func TestIA(t *testing.T) {
	tcs := []struct {
		a, x Value
	}{
		{False, False},
		{Unknown, True},
		{True, False},
	}

	for _, tc := range tcs {
		res := IA(tc.a)
		if res != tc.x {
			t.Errorf("(%v) want %v got %v", tc.a, tc.x, res)
		}
	}
}
