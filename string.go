package tinfer

// String is a string type.
type String struct {
	location string
}

// NewString creates a new string type.
func NewString(l string) String {
	return String{l}
}

// Accept accepts another type.
func (s String) Accept(t Type) error {
	if _, ok := t.(String); ok {
		return nil
	}

	return newInferenceError("not a string", t.Location())
}

// CanAccept checks if a type is acceptable.
func (String) CanAccept(t Type) bool {
	_, ok := t.(String)
	return ok
}

// Location returns where the type is defined.
func (s String) Location() string {
	return s.location
}
