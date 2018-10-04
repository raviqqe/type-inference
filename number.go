package tinfer

// Number is a number type.
type Number struct {
	location string
}

// NewNumber creates a new number type.
func NewNumber(l string) Number {
	return Number{l}
}

// Unify unifies 2 types.
func (n Number) Unify(t, tt *Type) error {
	if _, ok := (*tt).(Number); ok {
		return nil
	}

	return newInferenceError("not a number", (*tt).Location())
}

// Location returns where the type is defined.
func (n Number) Location() string {
	return n.location
}
