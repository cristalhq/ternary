package ternary

// AndB returns the result of Bochvar's logical conjunction on two values.
//
// +--------+-----------|
// |        |     B     |
// | A ∧ B  |---+---+---|
// |        | F | U | T |
// |----+---+---+---+---|
// |    | F | F | U | F |
// | A  | U | U | U | U |
// |    | T | F | U | T |
// +----+---+---+---+---+
//
func AndB(a Value, b Value) Value {
	switch {
	case a == True && b == True:
		return True
	case a == Unknown || b == Unknown:
		return Unknown
	default:
		return False
	}
}

// OrB returns the result of Bochvar's logical disjunction on two values.
//
// +--------+-----------+
// |        |     B     |
// | A ∨ B  |---+---+---|
// |        | F | U | T |
// |----+---+---+---+---|
// |    | F | F | U | T |
// | A  | U | U | U | U |
// |    | T | T | U | T |
// +----+---+---+---+---+
//
func OrB(a Value, b Value) Value {
	switch {
	case a == False && b == False:
		return False
	case a == Unknown || b == Unknown:
		return Unknown
	default:
		return True
	}
}

// ImpB returns the result of Bochvar's logical implication that is represented as "a implies b".
//
// +--------+-----------+
// |        |     B     |
// | A → B  |---+---+---|
// |        | F | U | T |
// |----+---+---+---+---|
// |    | F | T | U | T |
// | A  | U | U | U | T |
// |    | T | F | U | T |
// +----+---+---+---+---+
//
func ImpB(a Value, b Value) Value {
	switch {
	case b == True:
		return True
	case b == Unknown:
		return Unknown
	default:
		return Not(a)
	}
}
