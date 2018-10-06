package tinfer

// String is a number type.
type String struct {
	location string
}

// NewString creates a new number type.
func NewString(l string) String {
	return String{l}
}

// Unify unifies 2 types.
func (s String) Unify(t Type) error {
	if _, ok := t.(String); ok {
		return nil
	}

	return fallback(s, t, "not a number")
}

// Location returns where the type is defined.
func (s String) Location() string {
	return s.location
}
