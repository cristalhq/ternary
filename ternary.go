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

// Not returns the result of a logical negation of a value.
//
// +---+----+
// | A | ¬A |
// |---+----|
// | F |  T |
// | U |  U |
// | T |  F |
// +---+----+
//
func Not(a Value) Value {
	switch a {
	case False:
		return True
	case True:
		return False
	default:
		return Unknown
	}
}

// And returns the result of logical conjunction of two values.
//
// +--------+-----------|
// |        |     B     |
// | A ∧ B  |---+---+---|
// |        | F | U | T |
// |----+---+---+---+---|
// |    | F | F | F | F |
// | A  | U | F | U | U |
// |    | T | F | U | T |
// +----+---+---+---+---+
//
func And(a Value, b Value) Value {
	switch {
	case a == False || b == False:
		return False
	case a == Unknown || b == Unknown:
		return Unknown
	default:
		return True
	}
}

// Or returns the result of logical disjunction of two values.
//
// +--------+-----------+
// |        |     B     |
// | A ∨ B  |---+---+---|
// |        | F | U | T |
// |----+---+---+---+---|
// |    | F | F | U | T |
// | A  | U | U | U | T |
// |    | T | T | T | T |
// +----+---+---+---+---+
//
func Or(a Value, b Value) Value {
	switch {
	case a == True || b == True:
		return True
	case a == Unknown || b == Unknown:
		return Unknown
	default:
		return False
	}
}

// Imp Returns the result of logical implication that is represented as "a implies b".
// Same as OR(NOT(A), B).
//
// +--------+-----------+
// |        |     B     |
// | A → B  |---+---+---|
// |        | F | U | T |
// |----+---+---+---+---|
// |    | F | T | T | T |
// | A  | U | U | U | T |
// |    | T | F | U | T |
// +----+---+---+---+---+
//
func Imp(a Value, b Value) Value {
	return Or(Not(a), b)
}

// ImpL returns the result of Lukasiewicz's logical implication that is represented as "a implies b".
// Same as OR(NOT(A), B) except for both Unknown which is True.
//
// +--------+-----------+
// |        |     B     |
// | A → B  |---+---+---|
// |        | F | U | T |
// |----+---+---+---+---|
// |    | F | T | T | T |
// | A  | U | U | T | T |
// |    | T | F | U | T |
// +----+---+---+---+---+
//
func ImpL(a, b Value) Value {
	if a == Unknown && b == Unknown {
		return True
	}
	return Or(Not(a), b)
}

// Eqv Returns the result of logical biconditional of two values.
// Same as OR(AND(A, B), AND(NOT(A), NOT(B))).
//
// +--------+-----------+
// |        |     B     |
// | A ↔ B  |---+---+---|
// |        | F | U | T |
// |----+---+---+---+---|
// |    | F | T | U | F |
// | A  | U | U | U | U |
// |    | T | F | U | T |
// +----+---+---+---+---+
//
func Eqv(a Value, b Value) Value {
	if a == Unknown || b == Unknown {
		return Unknown
	}
	return FromBool(a == b)
}
