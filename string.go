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

	return fallback(s, t, "not a string")
}

// Location returns where the type is defined.
func (s String) Location() string {
	return s.location
}
