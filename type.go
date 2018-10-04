package tinfer

// Type is a type.
type Type interface {
	Unify(*Type, *Type) error
	Location() string
}
