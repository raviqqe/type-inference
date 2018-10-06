package tinfer

// Null is a null type.
type Null struct {
	location string
}

// NewNull creates a new null type.
func NewNull(l string) Null {
	return Null{l}
}

// Unify unifies 2 types.
func (s Null) Unify(t Type) error {
	if _, ok := t.(Null); ok {
		return nil
	}

	return fallback(s, t, "not a null")
}

// Location returns where the type is defined.
func (s Null) Location() string {
	return s.location
}
