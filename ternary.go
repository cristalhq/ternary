package ternary

// Value represents a ternary logic value.
type Value int8

const (
	// False represents a ternary false value.
	False Value = -1

	// Unknown  represents a ternary unknown value.
	Unknown Value = 0

	// True  represents a ternary true value.
	True Value = 1
)

// String returns string representation of the value.
func (value Value) String() string {
	switch value {
	case False:
		return "False"
	case Unknown:
		return "Unknown"
	case True:
		return "True"
	default:
		panic("unknown value")
	}
}

// Int returns integer representation of the value.
func (value Value) Int() int {
	switch value {
	case False:
		return -1
	case Unknown:
		return 0
	case True:
		return 1
	default:
		panic("unknown value")
	}
}

// AsBool returns true if the value is True, false otherwise.
func (value Value) AsBool() bool {
	return value == True
}

// IsTrue returns true if the value is True, false otherwise.
func (value Value) IsTrue() bool {
	return value == True
}

// IsUnknown returns true if the value is Unknown, false otherwise.
func (value Value) IsUnknown() bool {
	return value == Unknown
}

// IsFalse returns true if the value is False, false otherwise.
func (value Value) IsFalse() bool {
	return value == False
}

// Equal returns True if values are equal, False otherwise.
func Equal(a, b Value) Value {
	if a == b {
		return True
	}
	return False
}
