package tinfer

// Type is a type.
type Type interface {
	Accept(Type) error
	CanAccept(Type) bool
	Location() string
}
