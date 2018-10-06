package tinfer

// Binary is a binary type.
type Binary struct {
	location string
}

// NewBinary creates a new binary type.
func NewBinary(l string) Binary {
	return Binary{l}
}

// Unify unifies 2 types.
func (b Binary) Unify(t Type) error {
	if _, ok := t.(Binary); ok {
		return nil
	}

	return fallback(b, t, "not a binary")
}

// Location returns where the type is defined.
func (b Binary) Location() string {
	return b.location
}
