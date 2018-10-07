package tinfer

// Binary is a binary type.
type Binary struct {
	location string
}

// NewBinary creates a new binary type.
func NewBinary(l string) Binary {
	return Binary{l}
}

// Accept accepts another type.
func (b Binary) Accept(t Type) error {
	if _, ok := t.(Binary); ok {
		return nil
	}

	return newInferenceError("not a binary", t.Location())
}

// CanAccept checks if a type is acceptable.
func (Binary) CanAccept(t Type) bool {
	_, ok := t.(Binary)
	return ok
}

// Location returns where the type is defined.
func (b Binary) Location() string {
	return b.location
}
