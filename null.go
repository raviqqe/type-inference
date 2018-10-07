package tinfer

// Null is a null type.
type Null struct {
	location string
}

// NewNull creates a new null type.
func NewNull(l string) Null {
	return Null{l}
}

// Accept accepts another type.
func (n Null) Accept(t Type) error {
	if _, ok := t.(Null); ok {
		return nil
	}

	return newInferenceError("not a null", t.Location())
}

// CanAccept checks if a type is acceptable.
func (Null) CanAccept(t Type) bool {
	_, ok := t.(Null)
	return ok
}

// Location returns where the type is defined.
func (n Null) Location() string {
	return n.location
}
