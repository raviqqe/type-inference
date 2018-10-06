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
func (s Binary) Unify(t Type) error {
	if _, ok := t.(Binary); ok {
		return nil
	}

	return fallback(s, t, "not a binary")
}

// Location returns where the type is defined.
func (s Binary) Location() string {
	return s.location
}
