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
func (String) Unify(t, tt *Type) error {
	if _, ok := (*tt).(String); ok {
		return nil
	}

	return fallback(t, tt, "not a number")
}

// Location returns where the type is defined.
func (s String) Location() string {
	return s.location
}
