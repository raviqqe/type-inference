package tinfer

// Type is a type.
type Type interface {
	Unify(Type) error
	Location() string
}
