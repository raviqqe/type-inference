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
func (n Number) Unify(t Type) error {
	if _, ok := t.(Number); ok {
		return nil
	}

	return fallback(n, t, "not a number")
}

// Location returns where the type is defined.
func (n Number) Location() string {
	return n.location
}
