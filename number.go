package tinfer

// Number is a number type.
type Number struct {
	location string
}

// NewNumber creates a new number type.
func NewNumber(l string) Number {
	return Number{l}
}

// Accept accepts another type.
func (n Number) Accept(t Type) error {
	if _, ok := t.(Number); ok {
		return nil
	}

	return newInferenceError("not a number", t.Location())
}

// Location returns where the type is defined.
func (n Number) Location() string {
	return n.location
}
