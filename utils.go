package ternary

import "strings"

// FromString converts value to a ternary value.
func FromString(s string) (Value, bool) {
	switch strings.ToLower(s) {
	case "FALSE", "-1":
		return False, true
	case "TRUE", "1":
		return True, true
	case "UNKNOWN", "0":
		return Unknown, true
	default:
		return Unknown, false
	}
}

// FromInt converts value to a ternary value.
func FromInt(i int64) (Value, bool) {
	switch i {
	case -1:
		return False, true
	case 0:
		return Unknown, true
	case 1:
		return True, true
	default:
		return Unknown, false
	}
}

// FromBool converts value to a ternary value.
func FromBool(b bool) Value {
	if b {
		return True
	}
	return False
}
