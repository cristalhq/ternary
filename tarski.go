package ternary

// MA represents Tarski's "it is not false that...".
//
// +---+----+
// | A | MA |
// |---+----|
// | F |  F |
// | U |  T |
// | T |  T |
// +---+----+
//
func MA(a Value) Value {
	if a == False {
		return False
	}
	return True
}

// LA represents Tarski's "it is true that...".
//
// +---+----+
// | A | LA |
// |---+----|
// | F |  F |
// | U |  F |
// | T |  T |
// +---+----+
//
func LA(a Value) Value {
	if a == True {
		return True
	}
	return False
}

// IA represents Tarski "it is unknown that...".
//
// +---+----+
// | A | IA |
// |---+----|
// | F |  F |
// | U |  T |
// | T |  F |
// +---+----+
//
func IA(a Value) Value {
	if a == Unknown {
		return True
	}
	return False
}
